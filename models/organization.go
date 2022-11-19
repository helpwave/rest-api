package models

import "github.com/google/uuid"

type OrganizationBase struct {
	ID           uuid.UUID
	LongName     string `gorm:"default:NULL"`
	ShortName    string `gorm:"default:NULL"`
	AvatarUrl    string `gorm:"default:NULL"`
	ContactEmail string `gorm:"default:NULL"`
}

type Organization struct {
	OrganizationBase
	Users []User `gorm:"many2many:organizations_have_users"`
}
