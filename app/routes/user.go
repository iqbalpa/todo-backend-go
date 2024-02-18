package routes

import (
	"main/app/controller"
	"main/app/repository"
	"main/app/service"
	"main/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	userService service.UserService = service.NewUserService(repository.NewUserRepository())
)

func UserRoutes(api *gin.RouterGroup, db *gorm.DB){
	userController := controller.NewUserController(userService, db)

	user := api.Group("/user")
	{
		user.POST("/register", CreateUser(userController))
		user.POST("/login", LoginUser(userController))
	}
}

// ! REMOVE PASSWORD FROM THE PAYLOAD
func CreateUser(userController controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := userController.CreateUser(ctx)
		if err != nil {
			utils.FailedResponse(ctx, "FAILED", "failed to register", err)
		} else {
			utils.SuccessResponse(ctx, "register success", user)
		}
	}
}

func LoginUser(userController controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := userController.LoginUser(ctx)
		if err != nil {
			utils.FailedResponse(ctx, "FAILED", "failed to login", err)
		} else {
			utils.SuccessResponse(ctx, "login success", token)
		}
	}
}