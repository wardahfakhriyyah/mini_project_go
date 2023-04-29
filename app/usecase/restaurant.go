package usecase

import "miniproject_go_wardahfdn/app/model"

type RestaurantUseCase interface {
	CreateRestaurant(restaurant *model.Restaurant) (*model.Restaurant, error)
	GetRestaurantByID(restaurantID uint) (*model.Restaurant, error)
	UpdateRestaurant(restaurant *model.Restaurant) (*model.Restaurant, error)
	DeleteRestaurant(restaurantID uint) error
}
