package controller

import (
	"main/app/dto"
	"main/app/models"
	"main/app/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserController interface {
	CreateUser(ctx *gin.Context) (models.User, error)
	LoginUser(ctx *gin.Context) (string, error)
}

type userController struct {
	service service.UserService
	db *gorm.DB
}

func NewUserController(service service.UserService, db *gorm.DB) UserController {
	return &userController{
		service: service,
		db: db,
	}
} 

// ! REMOVE PASSWORD FROM THE PAYLOAD
func (uc *userController) CreateUser(ctx *gin.Context) (models.User, error) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return models.User{}, err
	}
	user, err = uc.service.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (uc *userController) LoginUser(ctx *gin.Context) (string, error) {
	var loginRequest dto.LoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		return "", err
	}
	// get the username and password
	username := loginRequest.Username
	password := loginRequest.Password
	token, err := uc.service.LoginUser(username, password)
	if err != nil {
		return "", err
	}
	return token, nil
}