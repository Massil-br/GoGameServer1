package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Model
	Username string `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Role     string `json:"role" `
}
