package HttpServer

import "github.com/aerosystems/stat-service/internal/models"

func (s *Server) setupRoutes() {
	s.echo.GET("/v1/events", s.eventHandler.GetEvents, s.AuthTokenMiddleware(models.CustomerRole))
}
