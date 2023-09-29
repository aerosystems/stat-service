package models

import "github.com/aerosystems/stat-service/internal/pagination"

type Event struct {
	Name         string `json:"event"`
	RawData      string `json:"rawData"`
	Domain       string `json:"domain"`
	Type         string `json:"type"`
	ProjectToken string `json:"projectToken"`
	Duration     int    `json:"duration"`
}

type EventRepository interface {
	GetByProjectToken(projectToken, eventType string, pagination pagination.Range) ([]Event, error)
}
