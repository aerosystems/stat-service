package models

import "github.com/google/uuid"

type Project struct {
	Id       int
	UserUuid uuid.UUID
	Name     string
	Token    string
}
