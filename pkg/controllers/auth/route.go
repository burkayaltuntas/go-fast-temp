package auth

import (
	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(r *gin.RouterGroup, authController *AuthController) {
	r.POST("/login", authController.Login)
}
