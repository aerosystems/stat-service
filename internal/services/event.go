package services

import (
	"github.com/aerosystems/stat-service/internal/models"
	RPCServices "github.com/aerosystems/stat-service/internal/rpc_services"
	RangeService "github.com/aerosystems/stat-service/pkg/range_service"
	"github.com/google/uuid"
)

type EventService interface {
	IsAccess(projectToken string, userUuid uuid.UUID) bool
	GetByProjectToken(projectToken string, kindEvent models.KindEvent, timeRange RangeService.TimeRange, pagination RangeService.LimitPagination) ([]models.Event, int, error)
}

type EventServiceImpl struct {
	projectRPC *RPCServices.ProjectRPC
	eventRepo  models.EventRepository
}

func NewEventServiceImpl(projectRPC *RPCServices.ProjectRPC, eventRepo models.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{
		projectRPC: projectRPC,
		eventRepo:  eventRepo,
	}
}

func (es *EventServiceImpl) IsAccess(projectToken string, userUuid uuid.UUID) bool {
	projectList, err := es.projectRPC.GetProjectList(userUuid)
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

func (es *EventServiceImpl) GetByProjectToken(projectToken string, kindEvent models.KindEvent, timeRange RangeService.TimeRange, pagination RangeService.LimitPagination) ([]models.Event, int, error) {
	return es.eventRepo.GetByProjectToken(projectToken, kindEvent, timeRange, pagination)
}
