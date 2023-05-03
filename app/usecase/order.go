package usecase

import "miniproject_go_wardahfdn/app/model"

type OrderUseCase interface {
	CreateOrder(order *model.Order) error
	GetAllOrder() ([]*model.Order, error)
	GetOrderByID(orderID uint) (*model.Order, error)
	UpdateOrder(order *model.Order) error
	DeleteOrder(orderID uint) error
}
