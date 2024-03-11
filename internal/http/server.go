package HttpServer

import (
	"fmt"
	"github.com/aerosystems/stat-service/internal/infrastructure/rest"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const webPort = 80

type Server struct {
	log          *logrus.Logger
	echo         *echo.Echo
	eventHandler *rest.EventHandler
	tokenService TokenService
}

func NewServer(
	log *logrus.Logger,
	eventHandler *rest.EventHandler,
	tokenService TokenService,

) *Server {
	return &Server{
		log:          log,
		echo:         echo.New(),
		eventHandler: eventHandler,
		tokenService: tokenService,
	}
}

func (s *Server) Run() error {
	s.setupMiddleware()
	s.setupRoutes()
	s.log.Infof("starting HTTP server stat-service on port %d\n", webPort)
	return s.echo.Start(fmt.Sprintf(":%d", webPort))
}
