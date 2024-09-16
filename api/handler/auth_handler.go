package handler

import (
	"api_gateway/config"
	"api_gateway/genproto/auth"
	"api_gateway/messagebroker"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authService auth.AuthServiceClient
	rabbitmq    *messagebroker.RabbitMQ
	cnf         *config.Config
}

func NewAuthHandler(auth auth.AuthServiceClient, cnf *config.Config) *AuthHandler {
	return &AuthHandler{
		authService: auth,
		cnf:         cnf,
	}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *AuthHandler) Register(c *gin.Context) {
	checkReq := RegisterRequest{}
	if err := c.BindJSON(&checkReq); err != nil {
		log.Println("Binding json failed: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't get user data from email"})
		return
	}

	checkRes, err := a.authService.CheckByEmail(context.Background(), &auth.CheckByEmailRequest{Email: checkReq.Email})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something wend wrong!"})
		return
	}

	if checkRes != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Email exists"})
		return
	}
	newToken := uuid.NewString()

	verificationLink := fmt.Sprintf("http://%s:%s/auth/verify-email?token=%s", a.cnf.Server.Host, a.cnf.Server.Port, newToken)

	a.authService.Register(context.Background(), &auth.RegisterRequest{
		Email:            checkReq.Email,
		Password:         checkReq.Password,
		VerificationLink: verificationLink,
	})

}

func (a *AuthHandler) Login(c *gin.Context) {

}
