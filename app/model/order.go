package model

import (
	"time"
)

type Order struct {
	OrderID      uint      `gorm:"primaryKey"`
	UserID       uint      `json:"-"`
	RestaurantID uint      `json:"-"`
	MenuID       uint      `json:"-"`
	Quantity     uint      `json:"quantity"`
	TotalPrice   uint      `json:"total_price"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
