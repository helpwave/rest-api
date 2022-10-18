package models

import (
    "github.com/google/uuid"
)

type EmergencyRoomBase struct {
    Name               string `binding:"required"`
    Location           Point  `binding:"required"`
    DisplayableAddress string `binding:"required"`
    Open               bool   `gorm:"column:is_open;default:true"`
    Utilization        int16  `gorm:"default:1"`
}

type EmergencyRoom struct {
    ID uuid.UUID
    EmergencyRoomBase
    Departments []Department `gorm:"many2many:rooms_have_departments" json:"-"`
}
