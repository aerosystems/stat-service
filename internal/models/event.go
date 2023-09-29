package models

type Event struct {
	Name         string `json:"event"`
	RawData      string `json:"rawData"`
	Domain       string `json:"domain"`
	Type         string `json:"type"`
	ProjectToken string `json:"projectToken"`
	Duration     int    `json:"duration"`
}

type EventRepository interface {
	GetByProjectToken(projectToken string, size, from int) ([]Event, error)
}
