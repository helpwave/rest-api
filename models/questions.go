package models

import "github.com/google/uuid"

type Question struct {
	ID       uuid.UUID `json:"id"`
	Question string    `json:"question" example:"Are children involved?"`
}

type Answer struct {
	ID         uuid.UUID `json:"id"`
	Answer     string    `json:"answer" example:"Yes"`
	Statement  string    `json:"statement" example:"There are children involved."`
	QuestionID uuid.UUID `json:"questionID"`
	Question   Question  `json:"question" gorm:"foreignKey:QuestionID"`
}
