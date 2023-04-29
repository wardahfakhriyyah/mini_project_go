package usecase

import "miniproject_go_wardahfdn/app/model"

type PaymentUseCase interface {
	CreatePayment(payment *model.Payment) (*model.Payment, error)
	GetInvoiceByID(paymentID uint) (*model.Payment, error)
	UpdateInvoice(payment *model.Payment) (*model.Payment, error)
	DeleteInvoice(paymentID uint) error
}
