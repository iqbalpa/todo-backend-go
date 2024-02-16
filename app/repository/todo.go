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

// Get todo instance by id
func (tr *TodoRepository) GetTodoById(id string) (models.Todo, error) {
	var todo models.Todo
	err := utils.DB.Where("id = ?", id).Take(&todo).Error
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

// Get all todo instances
func (tr *TodoRepository) GetTodos() ([]models.Todo, error){
	var todos []models.Todo
	err := utils.DB.Find(&todos).Error
	if err != nil {
		return []models.Todo{}, nil
	}
	return todos, nil
}