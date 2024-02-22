package service

import (
	"fmt"
	"main/app/models"
	"main/app/repository"
)

type TodoService interface {
	CreateTodo(models.Todo, int) (models.Todo, error)
	GetTodoById(string, int) (models.Todo, error)
	GetTodos(int) ([]models.Todo, error)
	DeleteTodoById(string, int) (string, error)
	UpdateTodoById(string, models.Todo, int) (models.Todo, error)
	FinishTodoById(string, int) (models.Todo, error)
}

type todoService struct {
	todoRepository *repository.TodoRepository
}

func NewTodoService(todoRepository *repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}


func (ts *todoService) CreateTodo(todo models.Todo, userId int) (models.Todo, error) {
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
	t.UserID = userId

	_, err := ts.todoRepository.CreateTodo(&t)
	if err != nil {
		return models.Todo{}, err
	}
	return t, nil
}

func (ts *todoService) GetTodoById(id string, userId int) (models.Todo, error) {
	todo, err := ts.todoRepository.GetTodoById(id, userId)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (ts *todoService) GetTodos(userId int) ([]models.Todo, error) {
	todos, err := ts.todoRepository.GetTodos(userId)
	if err != nil {
		return []models.Todo{}, err
	}
	return todos, nil
}

func (ts *todoService) DeleteTodoById(id string, userId int) (string, error) {
	str, err := ts.todoRepository.DeleteTodoById(id, userId)
	return str, err
}

func (ts *todoService) UpdateTodoById(id string, todo models.Todo, userId int) (models.Todo, error) {
	updatedTodo, err := ts.todoRepository.UpdateTodoById(id, &todo, userId)
	if err != nil {
		return models.Todo{}, err
	}
	return *updatedTodo, nil
}

func (ts *todoService) FinishTodoById(id string, userId int) (models.Todo, error) {
	todo, err := ts.todoRepository.FinishTodoById(id, userId)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}