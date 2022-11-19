package models

import (
	"github.com/google/uuid"
)

type EmergencyRoomBase struct {
	Name               string `json:"name" binding:"required" example:"Uniklinikum Münster"`
	Location           Point  `json:"location" binding:"required"`
	DisplayableAddress string `json:"displayableAddress" binding:"required" example:"Kardinal-von-Galen-Ring 10, 48149 Münster, Germany"`
	Open               bool   `json:"open" gorm:"column:is_open;default:true"`
	Utilization        int16  `json:"utilization" gorm:"default:1" example:"4"`
}

type EmergencyRoom struct {
	ID uuid.UUID `json:"id"`
	EmergencyRoomBase
	OrganizationID uuid.UUID    `json:"organizationID"`
	Organization   Organization `json:"organization"`
	Departments    []Department `json:"-" gorm:"many2many:rooms_have_departments" `
}
