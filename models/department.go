package models

import "github.com/google/uuid"

type Department struct {
	ID    uuid.UUID
	Name  string
	Rooms []EmergencyRoom `gorm:"many2many:rooms_have_departments"`
}
