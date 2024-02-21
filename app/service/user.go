package service

import (
	"fmt"
	"main/app/models"
	"main/app/repository"
)

type UserService interface {
	CreateUser(models.User) (models.User, error)
	LoginUser(string, string) (string, error)
	GetUserById(int) (models.User, error)
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

func (us *userService) LoginUser(username, password string) (string, error) {
	token, err := us.userRepository.LoginUser(username, password)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (us *userService) GetUserById(userId int) (models.User, error) {
	user, err := us.userRepository.GetUserById(userId)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}