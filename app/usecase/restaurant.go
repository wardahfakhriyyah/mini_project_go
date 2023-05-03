package usecase

import "miniproject_go_wardahfdn/app/model"

type RestaurantUseCase interface {
	CreateRestaurant(restaurant *model.Restaurant) error
	GetAllRestaurant() ([]*model.Restaurant, error)
	GetRestaurantByID(restaurantID uint) (*model.Restaurant, error)
	UpdateRestaurant(restaurant *model.Restaurant) error
	DeleteRestaurant(restaurantID uint) error
}
