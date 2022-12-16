package routes

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func InitRouter() (*gin.Engine, *gin.RouterGroup, *gin.RouterGroup) {
	r := gin.Default()

	corsSettings := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
	})

	r.Use(corsSettings)

	publicApi := r.Group("/api/v1/auth")
	{
		publicApi.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "publicly accessible",
			})
		})
	}

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "alive",
			})
		})
	}
	return r, apiv1, publicApi
}
