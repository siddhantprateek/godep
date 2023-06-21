package routes

import (
	config "godep/config"
	handler "godep/handler"

	"github.com/labstack/echo/v4"
)

func serverRoutes(router *echo.Echo) {

	router.GET("/", handler.Init)
	router.GET("/wokring", handler.Health)
}

func ApiServer() error {
	e := echo.New()

	serverRoutes(e)

	SERVER_PORT := config.GetEnv("PORT")
	if SERVER_PORT == "" {
		SERVER_PORT = "8080"
	}

	err := e.Start(":" + SERVER_PORT)
	return err
}
