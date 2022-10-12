package models

import "github.com/google/uuid"

type Question struct {
	ID       uuid.UUID
	Question string
}

type Answer struct {
	ID         uuid.UUID
	Answer     string
	Statement  string
	QuestionID uuid.UUID
	Question   Question `gorm:"foreignKey:QuestionID"`
}
