package models

import "github.com/google/uuid"

type Question struct {
	ID       uuid.UUID
	Question string `example:"Are children involved?"`
}

type Answer struct {
	ID         uuid.UUID
	Answer     string `example:"Yes"`
	Statement  string `example:"There are children involved."`
	QuestionID uuid.UUID
	Question   Question `gorm:"foreignKey:QuestionID"`
}
