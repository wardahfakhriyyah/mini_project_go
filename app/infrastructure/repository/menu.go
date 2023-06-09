package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"miniproject_go_wardahfdn/app/model"
)

type MenuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{db}
}

func (m *MenuRepository) CreateMenu(menu *model.Menu) error {
	menu.CreatedAt = time.Now()
	menu.UpdatedAt = time.Now()
	result := m.db.Create(&menu)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *MenuRepository) GetAllMenu() ([]*model.Menu, error) {
	var menu []*model.Menu
	result := m.db.Find(&menu)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no menu found")
		}
		return nil, result.Error
	}
	return menu, nil
}

func (m *MenuRepository) GetMenuById(menuID uint) (*model.Menu, error) {
	var menu model.Menu
	result := m.db.First(&menu, menuID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("menu not found with id: %d", menuID)
		}
		return nil, result.Error
	}
	return &menu, nil
}

func (m *MenuRepository) UpdateMenu(menu *model.Menu) error {
	menu.UpdatedAt = time.Now()
	result := m.db.Save(&menu)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *MenuRepository) DeleteMenu(menuID uint) error {
	result := m.db.Delete(&model.Menu{}, menuID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
