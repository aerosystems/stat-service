package repository

import (
	"encoding/json"
	"github.com/aerosystems/stat-service/internal/models"
	"github.com/aerosystems/stat-service/internal/transformator"
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

func (e *EventRepo) GetByProjectToken(projectToken string, size, from int) ([]models.Event, error) {
	query := `{
				  "_source": ["@timestamp", "message", "container"],
				  "query": {
					"bool": {
					  "must": [
						{
						  "match": {
							"container.id": "checkmail-service.log"
						  }
						},
						{
						  "match": {
							"message": "{\"projectToken\":\"` + projectToken + `\"}"
						  }
						}
					  ]
					}
				  }
				}`
	res, err := e.es.Search(
		e.es.Search.WithBody(strings.NewReader(query)),
		e.es.Search.WithPretty(),
		e.es.Search.WithSize(size),
		e.es.Search.WithFrom(from),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
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
	return eventList, nil
}
