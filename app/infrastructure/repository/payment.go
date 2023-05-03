package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"miniproject_go_wardahfdn/app/model"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db}
}

func (p *PaymentRepository) CreatePayment(payment *model.Payment) error {
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()
	result := p.db.Create(&payment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *PaymentRepository) GetAllPayment() ([]*model.Payment, error) {
	var payment []*model.Payment
	result := p.db.Find(&payment)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no payment found")
		}
		return nil, result.Error
	}
	return payment, nil
}

func (p *PaymentRepository) GetPaymentByID(paymentID uint) (*model.Payment, error) {
	var payment model.Payment
	result := p.db.First(&payment, paymentID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("payment not found with id: %d", paymentID)
		}
		return nil, result.Error
	}
	return &payment, nil
}

func (p *PaymentRepository) UpdatePayment(payment *model.Payment) error {
	payment.UpdatedAt = time.Now()
	result := p.db.Save(&payment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *PaymentRepository) DeletePayment(paymentID uint) error {
	result := p.db.Delete(&model.Payment{}, paymentID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
