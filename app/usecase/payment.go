package usecase

import "miniproject_go_wardahfdn/app/model"

type PaymentUseCase interface {
	CreatePayment(payment *model.Payment) error
	GetAllPayment() ([]*model.Payment, error)
	GetPaymentByID(paymentID uint) (*model.Payment, error)
	UpdatePayment(payment *model.Payment) error
	DeletePayment(paymentID uint) error
}
