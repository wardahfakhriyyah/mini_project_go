package model

import (
	"time"
)

type Menu struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        uint      `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	RestaurantID uint      `json:"-"`
}
