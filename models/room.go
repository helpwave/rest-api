package models

import (
	"github.com/google/uuid"
)

type EmergencyRoomBase struct {
	Name               string `binding:"required" example:"Uniklinikum Münster"`
	Location           Point  `binding:"required"`
	DisplayableAddress string `binding:"required" example:"Kardinal-von-Galen-Ring 10, 48149 Münster, Germany"`
	Open               bool   `gorm:"column:is_open;default:true"`
	Utilization        int16  `gorm:"default:1" example:"4"`
}

type EmergencyRoom struct {
	ID uuid.UUID
	EmergencyRoomBase
	Departments []Department `gorm:"many2many:rooms_have_departments" json:"-"`
}
