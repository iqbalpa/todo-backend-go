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
	todoService service.TodoService = service.NewTodoService(repository.NewTodoRepository())
)

func TodoRoutes(api *gin.RouterGroup, db *gorm.DB){
	todoController := controller.NewTodoController(todoService, db)

	todo := api.Group("/todo")
	{
		todo.POST("/", CreateTodo(todoController))
	}
}

func CreateTodo(todoController controller.TodoController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todo, err := todoController.CreateTodo(ctx)
		if err != nil {
			utils.FailedResponse(ctx, "FAILED", "failed to create todo", err)
		} else {
			utils.SuccessResponse(ctx, "todo created", todo)
		}
	}
}