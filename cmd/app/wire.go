//go:build wireinject
// +build wireinject

package main

import (
	"github.com/aerosystems/stat-service/internal/config"
	"github.com/aerosystems/stat-service/internal/infrastructure/http"
	"github.com/aerosystems/stat-service/internal/infrastructure/http/handlers"
	"github.com/aerosystems/stat-service/internal/repository/elastic"
	"github.com/aerosystems/stat-service/internal/repository/rpc"
	"github.com/aerosystems/stat-service/internal/usecases"
	ElasticClient "github.com/aerosystems/stat-service/pkg/elastic_client"
	"github.com/aerosystems/stat-service/pkg/logger"
	RpcClient "github.com/aerosystems/stat-service/pkg/rpc_client"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

//go:generate wire
func InitApp() *App {
	panic(wire.Build(
		wire.Bind(new(handlers.EventUsecase), new(*usecases.EventUsecase)),
		wire.Bind(new(usecases.EventRepository), new(*elastic.EventRepo)),
		wire.Bind(new(usecases.ProjectRepository), new(*rpc.ProjectRepo)),
		ProvideApp,
		ProvideLogger,
		ProvideConfig,
		ProvideHttpServer,
		ProvideLogrusLogger,
		ProvideBaseHandler,
		ProvideEventHandler,
		ProvideEventUsecase,
		ProvideEventRepo,
		ProvideProjectRepo,
		ProvideRpcClient,
		ProvideElasticClient,
	))
}

func ProvideApp(log *logrus.Logger, cfg *config.Config, httpServer *HttpServer.Server) *App {
	return NewApp(log, cfg, httpServer)
}

func ProvideLogger() *logger.Logger {
	panic(wire.Build(logger.NewLogger))
}

func ProvideConfig() *config.Config {
	panic(wire.Build(config.NewConfig))
}

func ProvideHttpServer(log *logrus.Logger, cfg *config.Config, eventHandler *handlers.EventHandler) *HttpServer.Server {
	return HttpServer.NewServer(log, cfg.AccessSecret, eventHandler)
}

func ProvideLogrusLogger(log *logger.Logger) *logrus.Logger {
	return log.Logger
}

func ProvideBaseHandler(log *logrus.Logger, cfg *config.Config) *handlers.BaseHandler {
	return handlers.NewBaseHandler(log, cfg.Mode)
}

func ProvideEventHandler(baseHandler *handlers.BaseHandler, eventUsecase handlers.EventUsecase) *handlers.EventHandler {
	panic(wire.Build(handlers.NewEventHandler))
}

func ProvideEventUsecase(eventRepo usecases.EventRepository, projectRepo usecases.ProjectRepository) *usecases.EventUsecase {
	panic(wire.Build(usecases.NewEventUsecase))
}

func ProvideEventRepo(es *elasticsearch.Client) *elastic.EventRepo {
	panic(wire.Build(elastic.NewEventRepo))
}

func ProvideProjectRepo(rpcClient *RpcClient.ReconnectRpcClient) *rpc.ProjectRepo {
	panic(wire.Build(rpc.NewProjectRepo))
}

func ProvideRpcClient(cfg *config.Config) *RpcClient.ReconnectRpcClient {
	return RpcClient.NewClient("tcp", cfg.ProjectServiceRpcAddress)
}

func ProvideElasticClient(cfg *config.Config) *elasticsearch.Client {
	return ElasticClient.NewClient(cfg.ElasticHost, cfg.ElasticPassword, cfg.ElasticCrtPath)
}
