package models

import "github.com/google/uuid"

type DepartmentBase struct {
	ID   uuid.UUID
	Name string
}

type Department struct {
	DepartmentBase
	Rooms               []EmergencyRoom `gorm:"many2many:rooms_have_departments"`
	NeededInEmergencies []Emergency     `gorm:"many2many:emergencies_need_departments"`
}
