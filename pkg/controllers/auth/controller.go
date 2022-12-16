package auth

import (
	"net/http"

	"github.com/burkayaltuntas/go-fast-temp/pkg/dto"
	"github.com/burkayaltuntas/go-fast-temp/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService AuthRepository
}

func NewAuthController(authService AuthRepository, routes *gin.RouterGroup) {
	ctr := &AuthController{
		AuthService: authService,
	}
	AddAuthRoutes(routes, ctr)
}

func (ctrl *AuthController) Login(ctx *gin.Context) {
	loginData := &loginData{}

	if err := ctx.ShouldBindJSON(loginData); err != nil {
		utils.HandleError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := loginData.Validate(); err != nil {
		utils.HandleError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	u, t, err := ctrl.AuthService.Login(loginData.Email, loginData.Password)
	if err != nil {
		utils.HandleError(ctx, http.StatusUnauthorized, "Error when login")
		return
	}
	utils.HandleSuccess(ctx, struct {
		User        *dto.UserDto `json:"user"`
		AccessToken string       `json:"accessToken"`
	}{u, t})
}
