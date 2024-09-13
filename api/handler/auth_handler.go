package handler

import (
	"api_gateway/genproto/auth"
	"api_gateway/messagebroker"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService auth.AuthServiceClient
	rabbitmq    *messagebroker.RabbitMQ
}

func NewAuthHandler(auth auth.AuthServiceClient) *AuthHandler {
	return &AuthHandler{
		authService: auth,
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

	a.authService.Register(context.Background(), &auth.RegisterRequest{
		Email:    checkReq.Email,
		Password: checkReq.Password,
	})

}
