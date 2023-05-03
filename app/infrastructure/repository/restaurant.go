package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"miniproject_go_wardahfdn/app/model"
)

type RestaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) *RestaurantRepository {
	return &RestaurantRepository{db}
}

func (r *RestaurantRepository) CreateRestaurant(restaurant *model.Restaurant) error {
	restaurant.CreatedAt = time.Now()
	restaurant.UpdatedAt = time.Now()
	result := r.db.Create(&restaurant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RestaurantRepository) GetAllRestaurant() ([]*model.Restaurant, error) {
	var restaurant []*model.Restaurant
	result := r.db.Find(&restaurant)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no restaurant found")
		}
		return nil, result.Error
	}
	return restaurant, nil
}

func (r *RestaurantRepository) GetRestaurantByID(restaurantID uint) (*model.Restaurant, error) {
	var restaurant model.Restaurant
	result := r.db.First(&restaurant, restaurantID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("restaurant not found with id: %d", restaurantID)
		}
		return nil, result.Error
	}
	return &restaurant, nil
}

func (r *RestaurantRepository) UpdateRestaurant(restaurant *model.Restaurant) error {
	restaurant.UpdatedAt = time.Now()
	result := r.db.Save(&restaurant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RestaurantRepository) DeleteRestaurant(restaurantID uint) error {
	result := r.db.Delete(&model.Restaurant{}, restaurantID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
