package repository

import (
	"encoding/json"
	"github.com/aerosystems/stat-service/internal/models"
	"github.com/aerosystems/stat-service/internal/transformator"
	RangeService "github.com/aerosystems/stat-service/pkg/range_service"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strings"
)

type EventRepo struct {
	es *elasticsearch.Client
}

func NewEventRepo(es *elasticsearch.Client) *EventRepo {
	return &EventRepo{
		es: es,
	}
}

func (e *EventRepo) GetByProjectToken(projectToken string, kindEvent models.KindEvent, timeRange RangeService.TimeRange, pagination RangeService.LimitPagination) ([]models.Event, int, error) {
	query := `{
				  "query": {
					"bool": {
					  "must": [
						{
						  "term": {
							"container.id": "checkmail-service.log"
						  }
						},
						{
						  "term": {
							"projectToken": "` + projectToken + `"
						  }
						},
						{
						  "term": {
							"kindEvent": "` + kindEvent.String() + `"
						  }
						},
						{
						  "range": {
							"time": {
							  "gte": "` + timeRange.Start.Format("2006-01-02T15:04:05.000Z") + `",
							  "lte": "` + timeRange.End.Format("2006-01-02T15:04:05.000Z") + `"
							}
						  }
						}
					  ]
					}
				  }
				}`
	res, err := e.es.Search(
		e.es.Search.WithBody(strings.NewReader(query)),
		e.es.Search.WithPretty(),
		e.es.Search.WithSize(pagination.Limit),
		e.es.Search.WithFrom(pagination.Offset),
		e.es.Search.WithSource("@timestamp", "message"),
		e.es.Search.WithSort("@timestamp:desc"),
		e.es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, 0, err
	}

	var eventList []models.Event
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		message := hit.(map[string]interface{})["_source"].(map[string]interface{})["message"].(string)
		ESEvent := transformator.ElasticSearchRecord{}
		err := json.Unmarshal([]byte(message), &ESEvent)
		if err != nil {
			log.Println(err)
			continue
		}
		event := ESEvent.ToEventModel()
		eventList = append(eventList, event)
	}

	total := int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	return eventList, total, nil
}
