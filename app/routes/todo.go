package routes

import (
	"main/app/controller"
	"main/app/middleware"
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
		todo.Use(middleware.Authorize())
		// todo.Use(middleware.UserIdExtractor())
		todo.GET("/", GetTodos(todoController))
		todo.POST("/", CreateTodo(todoController))
		todo.GET("/:id", GetTodoById(todoController))
		todo.PATCH("/:id", UpdateTodoById(todoController))
		todo.DELETE("/:id", DeleteTodoById(todoController))
		todo.PUT("/finish/:id", FinishTodoById(todoController))
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

func GetTodoById(todoController controller.TodoController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todo, err := todoController.GetTodoById(ctx)
		if err != nil {
			utils.FailedResponse(ctx, "FAILED", "failed to find the todo", err)
		} else {
			utils.SuccessResponse(ctx, "todo fetched", todo)
		}
	}
}

func GetTodos(todoController controller.TodoController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todos, err := todoController.GetTodos(ctx)
		if err != nil {
			utils.FailedResponse(ctx, "FAILED", "failed to find the todo", err)
		} else {
			utils.SuccessResponse(ctx, "todo fetched", todos)
		}
	}
}

func DeleteTodoById(todoController controller.TodoController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		str, err := todoController.DeleteTodoById(ctx)
		if err != nil {
			utils.FailedResponse(ctx, "FAILED", str, err)
		} else {
			utils.SuccessResponse(ctx, str, str)
		}
	}
}

func UpdateTodoById(todoController controller.TodoController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todo, err := todoController.UpdateTodoById(ctx)
		if err != nil {
			utils.FailedResponse(ctx, "FAILED", "failed to update todo", err)
		} else {
			utils.SuccessResponse(ctx, "todo updated", todo)
		}
	}
}

func FinishTodoById(todoController controller.TodoController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todo, err := todoController.FinishTodoById(ctx)
		if err != nil {
			utils.FailedResponse(ctx, "FAILED", "failed to finish the todo", err)
		} else {
			utils.SuccessResponse(ctx, "todo finished", todo)
		}
	}
}