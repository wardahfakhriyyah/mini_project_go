package model

import (
	"time"
)

type Menu struct {
	MenuID       uint      `gorm:"primaryKey"`
	RestaurantID uint      `json:"-"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        uint      `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
