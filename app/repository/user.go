package repository

import (
	"fmt"
	"main/app/models"
	"main/app/utils"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Create user instance
func (tr *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	// hash the password
	var err error
	user.Password, err = utils.HashingPassword(user.Password)
	if err != nil {
		return &models.User{}, fmt.Errorf("failed to hash the password")
	}
	err = utils.DB.Create(user).Error
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}
