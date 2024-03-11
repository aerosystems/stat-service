package HttpServer

import "github.com/aerosystems/customer-service/internal/models"

func (s *Server) setupRoutes() {
	e.GET("/v1/events", app.baseHandler.GetEvents, app.oauthMiddleware.AuthTokenMiddleware(models.CustomerRole))
}
