package main

import (
	_ "github.com/aerosystems/stat-service/docs"
	"github.com/aerosystems/stat-service/internal/models"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (app *Config) NewRouter() *echo.Echo {
	e := echo.New()

	docsGroup := e.Group("/docs")
	docsGroup.Use(app.basicAuthMiddleware.BasicAuthMiddleware)
	docsGroup.GET("/*", echoSwagger.WrapHandler)

	e.GET("/v1/events", app.baseHandler.GetEvents, app.oauthMiddleware.AuthTokenMiddleware(models.CustomerRole))

	return e
}
