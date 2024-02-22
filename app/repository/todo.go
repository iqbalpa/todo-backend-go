package repository

import (
	"main/app/models"
	"main/app/utils"
)

type TodoRepository struct{}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

// TODO: create restriction, therefore user can only access their own todos

// Create todo instance
func (tr *TodoRepository) CreateTodo(todo *models.Todo) (*models.Todo, error) {
	err := utils.DB.Create(todo).Error
	if err != nil {
		return &models.Todo{}, err
	}
	return todo, nil
}

// Get todo instance by id
func (tr *TodoRepository) GetTodoById(id string, userId int) (models.Todo, error) {
	var todo models.Todo
	err := utils.DB.Where("id = ?", id).Take(&todo).Error
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

// Get all todo instances
func (tr *TodoRepository) GetTodos(userId int) ([]models.Todo, error){
	var todos []models.Todo
	err := utils.DB.Where("user_id = ?", userId).Find(&todos).Error
	if err != nil {
		return []models.Todo{}, nil
	}
	return todos, nil
}

// delete announcement by id
func (tr *TodoRepository) DeleteTodoById(id string, userId int) (string, error) {
	var todo models.Todo
	err := utils.DB.Where("id = ?", id).Delete(&todo).Error
	if err != nil {
		return "failed to delete", err
	}
	return "todo deleted", nil
}

// update todo (title and desc)
func (tr *TodoRepository) UpdateTodoById(id string, todo *models.Todo, userId int) (*models.Todo, error) {
	err := utils.DB.Model(&models.Todo{}).Where("id = ?", id).Updates(todo).Error
	if err != nil {
		return &models.Todo{}, err
	}
	return todo, nil
}

// finish todo
func (tr *TodoRepository) FinishTodoById(id string, userId int) (models.Todo, error) {
	var todo models.Todo
	err := utils.DB.First(&todo, "id = ?", id).Error
	if err != nil {
		return models.Todo{}, err
	}

	// Update the IsFinished field
	todo.IsFinished = true
	err = utils.DB.Save(&todo).Error
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}
