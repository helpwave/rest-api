package models

import "github.com/google/uuid"

type UserBase struct {
	ID        uuid.UUID
	Email     string
	FullName  string `gorm:"default:NULL"`
	AvatarUrl string `gorm:"default:NULL"`
}

type User struct {
	UserBase
	PwBcrypt      string
	IsAdmin       bool
	Organizations []Organization `gorm:"many2many:organizations_have_users"`
}
