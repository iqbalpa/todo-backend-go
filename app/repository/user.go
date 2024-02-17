package repository

import (
	"main/app/models"
	"main/app/utils"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Create user instance
func (tr *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	err := utils.DB.Create(user).Error
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}
