package models

import (
	RangeService "github.com/aerosystems/stat-service/pkg/range_service"
	"time"
)

type Event struct {
	Name         string    `json:"event"`
	RawData      string    `json:"rawData"`
	Domain       string    `json:"domain"`
	Type         string    `json:"type"`
	ProjectToken string    `json:"projectToken"`
	Duration     int       `json:"duration"`
	Time         time.Time `json:"time"`
}

type EventRepository interface {
	GetByProjectToken(projectToken, eventType string, timeRange RangeService.TimeRange, pagination RangeService.LimitPagination) ([]Event, int, error)
}
