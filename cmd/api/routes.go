package main

import (
	_ "github.com/aerosystems/stat-service/docs"
	"github.com/aerosystems/stat-service/internal/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (app *Config) NewRouter() *echo.Echo {
	e := echo.New()

	e.Use(echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values echoMiddleware.RequestLoggerValues) error {
			app.Log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))

	docsGroup := e.Group("/docs")
	docsGroup.Use(middleware.BasicAuthMiddleware)
	docsGroup.GET("/*", echoSwagger.WrapHandler)

	e.GET("/v1/events", app.BaseHandler.GetEvents, middleware.AuthTokenMiddleware())

	return e
}
