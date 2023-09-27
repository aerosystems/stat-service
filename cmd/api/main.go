package main

import (
	"fmt"
	"github.com/aerosystems/lookup-service/pkg/logger"
	"github.com/aerosystems/stat-service/internal/handlers"
	"github.com/aerosystems/stat-service/internal/repository"
	"github.com/aerosystems/stat-service/pkg/elastic"
	"os"
)

const webPort = 80

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
