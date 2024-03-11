package usecases

import (
	"github.com/aerosystems/stat-service/internal/models"
	RangeService "github.com/aerosystems/stat-service/pkg/range_service"
	"github.com/google/uuid"
)

type ProjectRepository interface {
	GetProjectList(userUuid uuid.UUID) ([]models.Project, error)
	GetProject(projectToken string) (*models.Project, error)
}

type EventRepository interface {
	GetByProjectToken(projectToken string, kindEvent models.KindEvent, timeRange RangeService.TimeRange, pagination RangeService.LimitPagination) ([]models.Event, int, error)
}
