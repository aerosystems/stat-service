package main

import (
	"fmt"
	"github.com/aerosystems/stat-service/internal/handlers"
	"github.com/aerosystems/stat-service/internal/middleware"
	"github.com/aerosystems/stat-service/internal/repository"
	RPCServices "github.com/aerosystems/stat-service/internal/rpc_services"
	"github.com/aerosystems/stat-service/internal/services"
	"github.com/aerosystems/stat-service/pkg/elastic"
	"github.com/aerosystems/stat-service/pkg/logger"
	RPCClient "github.com/aerosystems/stat-service/pkg/rpc_client"
	"os"
)

const webPort = 80

// @title Stat Service API
// @version 1.0.0
// @description A part of microservice infrastructure, who responsible for statistics events

// @contact.name Artem Kostenko
// @contact.url https://github.com/aerosystems

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Should contain Access JWT Token, with the Bearer started

// @host gw.verifire.dev/stat
// @schemes https
// @BasePath /
func main() {
	log := logger.NewLogger(os.Getenv("HOSTNAME"))

	elasticsearchClient := elastic.NewClient()
	eventRepo := repository.NewEventRepo(elasticsearchClient)

	projectClientRPC := RPCClient.NewClient("tcp", "project-service:5001")
	projectRPC := RPCServices.NewProjectRPC(projectClientRPC)

	eventService := services.NewEventServiceImpl(projectRPC, eventRepo)

	baseHandler := handlers.NewBaseHandler(os.Getenv("APP_ENV"), log.Logger, eventService)

	accessTokenService := services.NewAccessTokenServiceImpl(os.Getenv("ACCESS_SECRET"))

	oauthMiddleware := middleware.NewOAuthMiddlewareImpl(accessTokenService)
	basicAuthMiddleware := middleware.NewBasicAuthMiddlewareImpl(os.Getenv("BASIC_AUTH_DOCS_USERNAME"), os.Getenv("BASIC_AUTH_DOCS_PASSWORD"))

	app := NewConfig(baseHandler, oauthMiddleware, basicAuthMiddleware)
	e := app.NewRouter()
	middleware.AddLog(e, log.Logger)
	if err := e.Start(fmt.Sprintf(":%d", webPort)); err != nil {
		log.Fatal(err)
	}
}
