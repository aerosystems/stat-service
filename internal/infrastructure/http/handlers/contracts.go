package handlers

import (
	"github.com/aerosystems/stat-service/internal/models"
	RangeService "github.com/aerosystems/stat-service/pkg/range_service"
	"github.com/google/uuid"
)

type EventUsecase interface {
	IsAccess(projectToken string, userUuid uuid.UUID) bool
	GetByProjectToken(projectToken string, kindEvent models.KindEvent, timeRange RangeService.TimeRange, pagination RangeService.LimitPagination) ([]models.Event, int, error)
}
