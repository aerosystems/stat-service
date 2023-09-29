package main

import (
	"fmt"
	"github.com/aerosystems/stat-service/internal/handlers"
	"github.com/aerosystems/stat-service/internal/repository"
	"github.com/aerosystems/stat-service/pkg/elastic"
	"github.com/aerosystems/stat-service/pkg/logger"
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

// @host api.verifire.com
// @schemes https
// @BasePath /
func main() {
	log := logger.NewLogger(os.Getenv("HOSTNAME"))

	elasticsearchClient := elastic.NewClient()
	eventRepo := repository.NewEventRepo(elasticsearchClient)

	baseHandler := handlers.NewBaseHandler(eventRepo)

	app := NewConfig(log.Logger, baseHandler)

	e := app.NewRouter()
	if err := e.Start(fmt.Sprintf(":%d", webPort)); err != nil {
		log.Fatal(err)
	}
}
