package usecases

import (
	"github.com/aerosystems/stat-service/internal/models"
	RangeService "github.com/aerosystems/stat-service/pkg/range_service"
	"github.com/google/uuid"
)

type EventUsecase struct {
	projectRepo ProjectRepository
	eventRepo   EventRepository
}

func NewEventUsecase(projectRepo ProjectRepository, eventRepo EventRepository) *EventUsecase {
	return &EventUsecase{
		projectRepo: projectRepo,
		eventRepo:   eventRepo,
	}
}

func (eu EventUsecase) IsAccess(projectToken string, userUuid uuid.UUID) bool {
	projectList, err := eu.projectRepo.GetProjectList(userUuid)
	if err != nil {
		return false
	}
	for _, project := range projectList {
		if project.UserUuid == userUuid && project.Token == projectToken {
			return true
		}
	}
	return false
}

func (eu EventUsecase) GetByProjectToken(projectToken string, kindEvent models.KindEvent, timeRange RangeService.TimeRange, pagination RangeService.LimitPagination) ([]models.Event, int, error) {
	return eu.eventRepo.GetByProjectToken(projectToken, kindEvent, timeRange, pagination)
}
