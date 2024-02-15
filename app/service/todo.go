package service

import (
	"fmt"
	"main/app/models"
	"main/app/repository"
)

type TodoService interface {
	CreateTodo(models.Todo) (models.Todo, error)
}

type todoService struct {
	todoRepository *repository.TodoRepository
}

func NewTodoService(todoRepository *repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}


func (ts *todoService) CreateTodo(todo models.Todo) (models.Todo, error) {
	if todo.Title == "" {
		return models.Todo{}, fmt.Errorf("title cannot be empty")
	}
	if todo.Description == ""{
		return models.Todo{}, fmt.Errorf("description cannot be empty")
	}

	// create todo instance
	t := models.Todo{}
	t.Title = todo.Title
	t.Description = todo.Description

	_, err := ts.todoRepository.CreateTodo(&t)
	if err != nil {
		return models.Todo{}, err
	}
	return t, nil
}