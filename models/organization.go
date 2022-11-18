package models

import "github.com/google/uuid"

type OrganizationBase struct {
	ID           uuid.UUID
	LongName     string
	ShortName    string `gorm:"default:NULL"`
	AvatarUrl    string `gorm:"default:NULL"`
	ContactEmail string
}

type Organization struct {
	OrganizationBase
	Users []User `gorm:"many2many:organizations_have_users"`
}
