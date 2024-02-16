package service

import (
	"fmt"
	"main/app/models"
	"main/app/repository"
)

type TodoService interface {
	CreateTodo(models.Todo) (models.Todo, error)
	GetTodoById(string) (models.Todo, error)
	GetTodos() ([]models.Todo, error)
	DeleteTodoById(string) (string, error)
	UpdateTodoById(string, models.Todo) (models.Todo, error)
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
	t.IsFinished = false

	_, err := ts.todoRepository.CreateTodo(&t)
	if err != nil {
		return models.Todo{}, err
	}
	return t, nil
}

func (ts *todoService) GetTodoById(id string) (models.Todo, error) {
	todo, err := ts.todoRepository.GetTodoById(id)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (ts *todoService) GetTodos() ([]models.Todo, error) {
	todos, err := ts.todoRepository.GetTodos()
	if err != nil {
		return []models.Todo{}, nil
	}
	return todos, nil
}

func (ts *todoService) DeleteTodoById(id string) (string, error) {
	str, err := ts.todoRepository.DeleteTodoById(id)
	return str, err
}

func (ts *todoService) UpdateTodoById(id string, todo models.Todo) (models.Todo, error) {
	updatedTodo, err := ts.todoRepository.UpdateTodoById(id, &todo)
	if err != nil {
		return models.Todo{}, err
	}
	return *updatedTodo, nil
}