package usecase

import "miniproject_go_wardahfdn/app/model"

type UserUseCase interface {
	CreateUser(user *model.User) error
	GetAllUser() ([]*model.User, error)
	GetUserByID(userID uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(userID uint) error
}
