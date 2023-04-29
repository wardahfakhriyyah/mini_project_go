package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"miniproject_go_wardahfdn/app/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *model.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) GetById(id uint) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with id: %d", id)
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with email: %s", email)
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Update(user *model.User) error {
	user.UpdatedAt = time.Now()
	result := r.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) Delete(id uint) error {
	result := r.db.Delete(&model.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
