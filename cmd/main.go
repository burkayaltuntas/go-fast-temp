package main

import (
	"os"

	"github.com/burkayaltuntas/go-fast-temp/pkg/config"
	"github.com/burkayaltuntas/go-fast-temp/pkg/controllers/auth"
	"github.com/burkayaltuntas/go-fast-temp/pkg/controllers/home"
	"github.com/burkayaltuntas/go-fast-temp/pkg/data"
	"github.com/burkayaltuntas/go-fast-temp/pkg/middlewares"
	"github.com/burkayaltuntas/go-fast-temp/pkg/routes"
	authsrv "github.com/burkayaltuntas/go-fast-temp/pkg/services/auth"
	homesrv "github.com/burkayaltuntas/go-fast-temp/pkg/services/home"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	jwt_secret := os.Getenv("JWT_SECRET")
	dbConfig := config.GetDbConfig()
	db := data.Connect(dbConfig)
	app, routes, publicRoutes := routes.InitRouter()

	// injecting controller and service dependecies
	authService := authsrv.NewAuthService(db, jwt_secret)

	routes.Use(middlewares.AuthChecker(authService))
	routes.Use(middlewares.Logger())

	auth.NewAuthController(authService, publicRoutes)

	homeService := homesrv.NewHomeService(db)
	home.NewHomeController(homeService, routes)

	app.Run(":8080") // listen and serve on 0.0.0.0:8080
}
