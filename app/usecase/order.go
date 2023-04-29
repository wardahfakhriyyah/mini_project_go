package usecase

import "miniproject_go_wardahfdn/app/model"

type OrderUseCase interface {
	CreateOrder(order *model.Order) (*model.Order, error)
	GetOrderByID(orderID uint) (*model.Order, error)
	GetOrdersByUserID(userID uint) ([]*model.Order, error)
	GetOrdersByRestaurantID(restaurantID uint) ([]*model.Order, error)
	UpdateOrder(order *model.Order) (*model.Order, error)
	DeleteOrder(orderID uint) error
}
