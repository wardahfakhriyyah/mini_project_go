package usecase

import "miniproject_go_wardahfdn/app/model"

type MenuUseCase interface {
	CreateMenu(menu *model.Menu) error
	GetMenuByID(menuID uint) (*model.Menu, error)
	UpdateMenu(menu *model.Menu) error
	DeleteMenu(menuID uint) error
}
