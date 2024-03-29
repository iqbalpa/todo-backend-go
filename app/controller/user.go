package controller

import (
	"fmt"
	"log"
	"main/app/dto"
	"main/app/models"
	"main/app/service"
	"main/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserController interface {
	CreateUser(*gin.Context) (models.User, error)
	LoginUser(*gin.Context) (map[string]string, error)
	GetUserById(*gin.Context) (models.User, error)
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
	// remove the password from the payload
	user.Password = ""
	return user, nil
}

func (uc *userController) LoginUser(ctx *gin.Context) (map[string]string, error) {
	var loginRequest dto.LoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		return make(map[string]string), err
	}
	// get the username and password
	username := loginRequest.Username
	password := loginRequest.Password
	token, err := uc.service.LoginUser(username, password)
	if err != nil {
		return make(map[string]string), err
	}
	tokenMap := map[string]string{"token": token}
	return tokenMap, nil
}

func (uc *userController) GetUserById(ctx *gin.Context) (models.User, error) {
	// get the userId
	token := ctx.GetString("token")
	log.Println("controller/user.go\ntoken: ", token)
	userId, err := utils.ExtractClaimsUserId(token)
	log.Println("user id: ", userId)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get the userId")
	}
	user, err := uc.service.GetUserById(userId)
	if err != nil {
		return models.User{}, err
	}
	user.Password = ""
	return user, nil
}