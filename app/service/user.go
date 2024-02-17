package service

import (
	"fmt"
	"main/app/models"
	"main/app/repository"
)

type UserService interface {
	CreateUser(models.User) (models.User, error)
}

type userService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (us *userService) CreateUser(user models.User) (models.User, error) {
	if user.Username == "" {
		return models.User{}, fmt.Errorf("username cannot be empty")
	}
	if user.Password == "" {
		return models.User{}, fmt.Errorf("password cannot be empty")
	}
	if user.Name == "" {
		return models.User{}, fmt.Errorf("name cannot be empty")
	}
	u, err := us.userRepository.CreateUser(&user)
	if err != nil {
		return models.User{}, err
	}
	return *u, nil
}
