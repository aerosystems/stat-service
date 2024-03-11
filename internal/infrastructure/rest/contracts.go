package rest

import "github.com/aerosystems/stat-service/internal/models"

type EventUsecase interface {
	GetEvents() ([]*models.Event, error)
}
