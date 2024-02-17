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
		user.POST("/", CreateUser(userController))
	}
}

func CreateUser(userController controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := userController.CreateUser(ctx)
		if err != nil {
			utils.FailedResponse(ctx, "FAILED", "failed to create todo", err)
		} else {
			utils.SuccessResponse(ctx, "todo created", user)
		}
	}
}
