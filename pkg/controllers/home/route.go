package home

import (
	"github.com/gin-gonic/gin"
)

func AddHomeRoutes(r *gin.RouterGroup, homeController *HomeController) {

	apiv1 := r.Group("/home")
	{
		apiv1.GET("/city", homeController.GetCityList)

	}
}
