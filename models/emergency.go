package models

import (
	"github.com/google/uuid"
	"time"
)

type Emergency struct {
	ID                uuid.UUID
	StartLoc          Point
	TimeStamp         time.Time     `gorm:"default:now()"`
	EmergencyRoomID   uuid.UUID     `gorm:"default:NULL"` // explicitly set the FK to NULL, else PG is confused
	EmergencyRoom     EmergencyRoom `gorm:"foreignKey:EmergencyRoomID"`
	NeededDepartments []Department  `gorm:"many2many:emergencies_need_departments"`
	Answers           []Answer      `gorm:"many2many:emergency_related_answers"`
}
