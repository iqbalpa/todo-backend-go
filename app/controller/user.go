package controller

import (
	"main/app/models"
	"main/app/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserController interface {
	CreateUser(ctx *gin.Context) (models.User, error)
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
