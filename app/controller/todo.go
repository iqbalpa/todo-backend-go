package controller

import (
	"fmt"
	"main/app/models"
	"main/app/service"
	"main/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TodoController interface {
	CreateTodo(ctx *gin.Context) (models.Todo, error)
	GetTodoById(ctx *gin.Context) (models.Todo, error)
	GetTodos(ctx *gin.Context) ([]models.Todo, error)
	DeleteTodoById(ctx *gin.Context) (string, error)
	UpdateTodoById(ctx *gin.Context) (models.Todo, error)
	FinishTodoById(ctx *gin.Context) (models.Todo, error)
}

type todoController struct {
	service service.TodoService
	db *gorm.DB
}

func NewTodoController(service service.TodoService, db *gorm.DB) TodoController {
	return &todoController{
		service: service,
		db: db,
	}
} 

func (tc *todoController) CreateTodo(ctx *gin.Context) (models.Todo, error) {
	var todo models.Todo
	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		return models.Todo{}, err
	}
	// get the userId
	token := ctx.GetString("token")
	userId, err := utils.ExtractClaimsUserId(token)
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to get the userId")
	}
	// call the service function
	_, err = tc.service.CreateTodo(todo, userId)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (tc *todoController) GetTodoById(ctx *gin.Context) (models.Todo, error) {
	id := ctx.Param("id")
	todo, err := tc.service.GetTodoById(id)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (tc *todoController) GetTodos(ctx *gin.Context) ([]models.Todo, error) {
	todos, err := tc.service.GetTodos()
	if err != nil {
		return []models.Todo{}, err
	}
	return todos, nil
}

func (tc *todoController) DeleteTodoById(ctx *gin.Context) (string, error) {
	id := ctx.Param("id")
	str, err := tc.service.DeleteTodoById(id)
	return str, err
}

func (tc *todoController) UpdateTodoById(ctx *gin.Context) (models.Todo, error) {
	id := ctx.Param("id")
	var todo models.Todo
	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		return models.Todo{}, err
	}
	todo, err = tc.service.UpdateTodoById(id, todo)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (tc *todoController) FinishTodoById(ctx *gin.Context) (models.Todo, error) {
	id := ctx.Param("id")
	todo, err := tc.service.FinishTodoById(id)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}