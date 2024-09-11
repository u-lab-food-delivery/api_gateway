package api

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	authen := r.Group("/auth")
	{
		authen.POST("/register")
		authen.POST("/login")
		authen.POST("/verify-email")
		authen.POST("/logout")
	}
}
