package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
