package models

import "github.com/google/uuid"

const Admin = "admin"

type UserBase struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"fullName" gorm:"default:NULL"`
	AvatarUrl string    `json:"avatarUrl" gorm:"default:NULL"`
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
