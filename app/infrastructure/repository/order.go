package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"miniproject_go_wardahfdn/app/model"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (o *OrderRepository) CreateOrder(order *model.Order) error {
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	result := o.db.Create(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (o *OrderRepository) GetAllOrder() ([]*model.Order, error) {
	var order []*model.Order
	result := o.db.Find(&order)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no order found")
		}
		return nil, result.Error
	}
	return order, nil
}

func (o *OrderRepository) GetOrderById(orderID uint) (*model.Order, error) {
	var order model.Order
	result := o.db.First(&order, orderID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("order not found with id : %d", orderID)
		}
		return nil, result.Error
	}
	return &order, nil
}

func (o *OrderRepository) UpdateOrder(order *model.Order) error {
	order.UpdatedAt = time.Now()
	result := o.db.Save(&order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (o *OrderRepository) DeleteOrder(orderID uint) error {
	result := o.db.Delete(&model.Order{}, orderID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
