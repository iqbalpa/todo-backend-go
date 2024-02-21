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

// Create user instance (register)
func (ur *UserRepository) CreateUser(user *models.User) (*models.User, error) {
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

// login user
func (ur *UserRepository) LoginUser(username, password string) (string, error) {
	// retrieve user data
	var err error
	var userData models.User
	err = utils.DB.Model(&models.User{}).Where("username = ?", username).Take(&userData).Error
	if err != nil {
		return "", fmt.Errorf("user with username %s is not found", username)
	}
	// verify password
	err = utils.VerifyPassword(password, userData.Password)
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}
	// generate token
	var token string
	token, err = utils.CreateToken(userData)
	if err != nil {
		return "", err
	}
	return token, nil
}

// get user data
func (ur *UserRepository) GetUserById(userId int) (models.User, error) {
	var user models.User
	err := utils.DB.Preload("Todos").First(&user, userId).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}