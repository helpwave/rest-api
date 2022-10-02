package models

import (
	"github.com/google/uuid"
)

type EmergencyRoom struct {
	ID                 uuid.UUID
	Name               string
	Location           Point
	DisplayableAddress string
	Open               bool         `gorm:"column:is_open;default:true"`
	Utilization        int16        `gorm:"default:1"`
	Departments        []Department `gorm:"many2many:rooms_have_departments"`
}
