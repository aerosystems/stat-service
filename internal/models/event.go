package models

import (
	RangeService "github.com/aerosystems/stat-service/pkg/range_service"
	"time"
)

type Event struct {
	Name         string    `json:"name"`
	RawData      string    `json:"rawData"`
	ErrorCode    int       `json:"errorCode,omitempty"`
	ErrorMessage string    `json:"errorMessage,omitempty"`
	Domain       string    `json:"domain,omitempty"`
	DomainType   string    `json:"domainType,omitempty"`
	ProjectToken string    `json:"projectToken"`
	Duration     int       `json:"duration"`
	CreatedAt    time.Time `json:"createdAt"`
}

type EventRepository interface {
	GetByProjectToken(projectToken string, kindEvent KindEvent, timeRange RangeService.TimeRange, pagination RangeService.LimitPagination) ([]Event, int, error)
}
