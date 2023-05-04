package usecase

import "miniproject_go_wardahfdn/app/model"

type AuthUseCase interface {
	Register(user *model.User) (*model.User, error)
	Login(email, password string) (*model.User, error)
}
