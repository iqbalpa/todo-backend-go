package controller

import (
	"main/app/models"
	"main/app/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TodoController interface {
	CreateTodo(ctx *gin.Context) (models.Todo, error)
	GetTodoById(ctx *gin.Context) (models.Todo, error)
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
	_, err = tc.service.CreateTodo(todo)
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