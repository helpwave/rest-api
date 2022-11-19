package models

import (
	"time"

	"github.com/google/uuid"
)

type Emergency struct {
	ID                uuid.UUID     `json:"id"`
	StartLoc          Point         `json:"startLoc"`
	TimeStamp         time.Time     `json:"timeStamp" gorm:"default:now()"`
	EmergencyRoomID   uuid.UUID     `json:"emergencyRoomID" gorm:"default:NULL"` // explicitly set the FK to NULL, else PG is confused
	EmergencyRoom     EmergencyRoom `json:"emergencyRoom" gorm:"foreignKey:EmergencyRoomID"`
	NeededDepartments []Department  `json:"neededDepartments" gorm:"many2many:emergencies_need_departments"`
	Answers           []Answer      `json:"answers" gorm:"many2many:emergency_related_answers"`
}
