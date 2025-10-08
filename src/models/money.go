package models

import "github.com/google/uuid"

type Money struct {
	Model
	UserId uuid.UUID `json:"userId" gorm:"not null"`
	Amount float64   `json:"amount" gorm:"default:0"`
}
