package models

type Event struct {
	Name         string `json:"event"`
	RawData      string `json:"rawData"`
	Domain       string `json:"domain"`
	Type         string `json:"type"`
	ProjectToken string `json:"projectToken"`
	Duration     int64  `json:"duration"`
	Source       string `json:"source"`
}

type EventRepository interface {
	GetByProjectToken(projectToken string) ([]Event, error)
}
