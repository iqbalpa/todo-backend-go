package routes

import (
	"main/app/controller"
	"main/app/repository"
	"main/app/service"
	"net/http"

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
			ctx.JSON(http.StatusBadRequest, "failed to create")
		} else {
			ctx.JSON(http.StatusOK, todo)
		}
		return
	}
}