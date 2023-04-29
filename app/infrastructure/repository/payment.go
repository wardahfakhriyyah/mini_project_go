package repository

import (
	"errors"

	"gorm.io/gorm"

	"miniproject_go_wardahfdn/app/model"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db}
}

func (r *PaymentRepository) CreatePayment(payment *model.Payment) error {
	if err := r.db.Create(payment).Error; err != nil {
		return err
	}
	return nil
}

func (r *PaymentRepository) GetPaymentByID(paymentID uint) (*model.Payment, error) {
	var payment model.Payment
	if err := r.db.First(&payment, paymentID).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *PaymentRepository) UpdatePayment(payment *model.Payment) error {
	if payment.ID == 0 {
		return errors.New("invalid payment id")
	}
	if err := r.db.Save(payment).Error; err != nil {
		return err
	}
	return nil
}

func (r *PaymentRepository) DeletePayment(payment *model.Payment) error {
	if payment.ID == 0 {
		return errors.New("invalid payment id")
	}
	if err := r.db.Delete(payment).Error; err != nil {
		return err
	}
	return nil
}
