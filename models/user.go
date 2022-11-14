package models

import "github.com/google/uuid"

const Admin = "admin"

type UserBase struct {
	ID        uuid.UUID
	Email     string
	FullName  string `gorm:"default:NULL"`
	AvatarUrl string `gorm:"default:NULL"`
}

type User struct {
	UserBase
	PwBcrypt      string
	GlobalRoles   []GlobalRole
	Organizations []Organization `gorm:"many2many:organizations_have_users"`
}

type GlobalRole struct {
	UserID uuid.UUID
	Role   string
}
