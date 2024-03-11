package models

import (
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

type KindEvent string

const (
	InspectEvent KindEvent = "inspect"
)

func (k KindEvent) String() string {
	return string(k)
}
