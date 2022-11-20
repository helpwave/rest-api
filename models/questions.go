package models

import "github.com/google/uuid"

type Question struct {
	ID       uuid.UUID `json:"id"`
	Question string    `json:"question" gorm:"default:NULL" example:"Are children involved?"`
}

type Answer struct {
	ID         uuid.UUID `json:"id"`
	Answer     string    `json:"answer" gorm:"default:NULL" example:"Yes"`
	Statement  string    `json:"statement" gorm:"default:NULL" example:"There are children involved."`
	QuestionID uuid.UUID `json:"questionID" gorm:"default:NULL"`
	Question   Question  `json:"question" gorm:"foreignKey:QuestionID"`
}
