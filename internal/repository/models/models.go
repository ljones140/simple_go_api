package models

import "github.com/google/uuid"

type Object struct {
	ID   uuid.UUID
	Name string
}
