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

func (u *UserRepository) CreateUser(user *model.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	result := u.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) GetUserByID(userID uint) (*model.User, error) {
	var user model.User
	result := u.db.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with id: %d", userID)
		}
		return nil, result.Error
	}
	return &user, nil
}

// func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
// 	var user model.User
// 	result := r.db.Where("email = ?", email).First(&user)
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return nil, fmt.Errorf("user not found with email: %s", email)
// 		}
// 		return nil, result.Error
// 	}
// 	return &user, nil
// }

func (u *UserRepository) UpdateUser(user *model.User) error {
	user.UpdatedAt = time.Now()
	result := u.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) DeleteUser(userID uint) error {
	result := u.db.Delete(&model.User{}, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
