package model

import (
	"time"
)

type User struct {
	UserID    uint      `gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `gorm:"email" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
