package model

import (
	"time"
)

type Payment struct {
	PaymentID     uint      `gorm:"primaryKey"`
	OrderID       uint      `json:"-"`
	TotalAmount   uint      `json:"total_amount"`
	PaymentMethod string    `json:"method"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
