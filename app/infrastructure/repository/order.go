package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"miniproject_go_wardahfdn/app/model"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) Create(order *model.Order) error {
	result := r.db.Create(order)
	if result.Error != nil {
		return fmt.Errorf("error creating order: %v", result.Error)
	}
	return nil
}

func (r *OrderRepository) GetById(id uint) (*model.Order, error) {
	var order model.Order
	result := r.db.First(&order, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("order with ID %v not found", id)
	}
	if result.Error != nil {
		return nil, fmt.Errorf("error retrieving order: %v", result.Error)
	}
	return &order, nil
}

func (r *OrderRepository) Update(order *model.Order) error {
	result := r.db.Save(order)
	if result.Error != nil {
		return fmt.Errorf("error updating order: %v", result.Error)
	}
	return nil
}

func (r *OrderRepository) Delete(order *model.Order) error {
	result := r.db.Delete(order)
	if result.Error != nil {
		return fmt.Errorf("error deleting order: %v", result.Error)
	}
	return nil
}
