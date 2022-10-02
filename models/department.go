package models

import "github.com/google/uuid"

type Department struct {
	ID   uuid.UUID
	Name string
}
