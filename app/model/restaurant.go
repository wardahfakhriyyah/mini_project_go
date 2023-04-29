package model

import (
	"time"
)

type Restaurant struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
