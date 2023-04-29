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

func (m *MenuRepository) Create(menu *model.Menu) error {
	menu.CreatedAt = time.Now()
	menu.UpdatedAt = time.Now()
	result := m.db.Create(&menu)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *MenuRepository) GetById(id uint) (*model.Menu, error) {
	var menu model.Menu
	result := m.db.First(&menu, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("menu not found with id: %d", id)
		}
		return nil, result.Error
	}
	return &menu, nil
}

func (m *MenuRepository) Update(menu *model.Menu) error {
	menu.UpdatedAt = time.Now()
	result := m.db.Save(&menu)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *MenuRepository) Delete(id uint) error {
	result := m.db.Delete(&model.Menu{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
