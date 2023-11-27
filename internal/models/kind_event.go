package models

type KindEvent string

const (
	InspectEvent KindEvent = "inspect"
)

func (k KindEvent) String() string {
	return string(k)
}
