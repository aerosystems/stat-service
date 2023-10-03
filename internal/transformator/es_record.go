package transformator

import (
	"github.com/aerosystems/stat-service/internal/models"
	"time"
)

type ElasticSearchRecord struct {
	Domain       string `json:"domain"`
	Duration     int    `json:"duration"`
	Event        string `json:"event"`
	Level        string `json:"level"`
	Msg          string `json:"msg"`
	ProjectToken string `json:"projectToken"`
	RawData      string `json:"rawData"`
	Source       string `json:"source"`
	Time         string `json:"time"`
	Type         string `json:"type"`
}

func (es *ElasticSearchRecord) ToEventModel() models.Event {
	t, err := time.Parse(time.RFC3339, es.Time)
	if err != nil {
		t = time.Now()
	}
	return models.Event{
		Name:         es.Event,
		RawData:      es.RawData,
		Domain:       es.Domain,
		Type:         es.Type,
		ProjectToken: es.ProjectToken,
		Duration:     es.Duration,
		Time:         t,
	}
}
