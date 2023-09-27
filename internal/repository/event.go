package repository

import (
	"github.com/aerosystems/stat-service/internal/models"
	"github.com/elastic/go-elasticsearch/v8"
)

type EventRepo struct {
	es *elasticsearch.Client
}

func NewEventRepo(es *elasticsearch.Client) *EventRepo {
	return &EventRepo{
		es: es,
	}
}

func (e *EventRepo) GetByProjectToken(projectToken string) ([]models.Event, error) {
	return nil, nil
}
