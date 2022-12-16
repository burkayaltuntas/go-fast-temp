package home

import (
	"github.com/burkayaltuntas/go-fast-temp/pkg/utils"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	HomeService HomeRepository
}

func NewHomeController(homeService HomeRepository, routes *gin.RouterGroup) {
	ctr := &HomeController{
		HomeService: homeService,
	}
	AddHomeRoutes(routes, ctr)
}

func (c *HomeController) GetCityList(ctx *gin.Context) {
	projects := c.HomeService.GetCityList()
	utils.HandleSuccess(ctx, projects)
}
