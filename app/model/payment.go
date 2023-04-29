package model

import (
	"time"
)

type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	TotalAmount   uint      `json:"total_amount"`
	PaymentMethod string    `json:"method"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	OrderID       uint      `json:"-"`
}
