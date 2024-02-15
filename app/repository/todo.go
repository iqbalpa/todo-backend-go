package repository

import (
	"main/app/models"
	"main/app/utils"
)

type TodoRepository struct{}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

// Create todo instance
func (tr *TodoRepository) CreateTodo(todo *models.Todo) (*models.Todo, error) {
	err := utils.DB.Create(todo).Error
	if err != nil {
		return &models.Todo{}, err
	}
	return todo, nil
}