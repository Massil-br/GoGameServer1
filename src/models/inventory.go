package models

import "github.com/google/uuid"

type Inventory struct {
	Model
	UserId uuid.UUID `json:"userId" gorm:"not null"`
	Slots  [36]string
}
