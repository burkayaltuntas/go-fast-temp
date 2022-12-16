package middlewares

import (
	"net/http"
	"strings"

	"github.com/burkayaltuntas/go-fast-temp/pkg/common"
	"github.com/burkayaltuntas/go-fast-temp/pkg/services/auth"
	"github.com/burkayaltuntas/go-fast-temp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthChecker(srv *auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		t := extractToken(bearerToken)
		if t == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseModel{Status: "401", Message: "Unauthorized"})
			return
		}
		token, err := srv.ValidateToken(t)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		customClaims := token.Claims.(jwt.MapClaims)

		user := common.ContextUser{
			Email:   customClaims["email"].(string),
			Role:    common.Role((customClaims["role"].(float64))),
			Id:      customClaims["id"].(string),
			Name:    customClaims["name"].(string),
			Surname: customClaims["surname"].(string),
		}
		c.Set("user", user)
		c.Next()
	}
}

func extractToken(whole string) string {
	if len(strings.Split(whole, " ")) == 2 {
		return strings.Split(whole, " ")[1]
	}
	return ""
}
