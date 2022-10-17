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

// DepartmentsToBases creates an array of DepartmentsToBases for a given array of Departments
func DepartmentsToBases(deps []Department) []DepartmentBase {
	bases := make([]DepartmentBase, len(deps))
	for i := range deps {
		bases[i] = deps[i].DepartmentBase
	}
	return bases
}
