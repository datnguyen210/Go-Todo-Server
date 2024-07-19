package usecase

import (
	"datnguyen/todo/internal/entity"
)

type Todo = entity.Todo

type TodoRepository interface {
	IndexTodos() ([]Todo, error)
	ReadTodoByID(id int64) (Todo, error)
	CreateTodo(todo Todo) (Todo, error)
	UpdateTodo(todo Todo) (Todo, error)
	DeleteTodoByID(id int64) error
}

type TodoUseCase struct {
	TodoRepo TodoRepository
}

func NewTodoUseCase(todoRepo TodoRepository) *TodoUseCase {
	return &TodoUseCase{TodoRepo: todoRepo}
}

func (uc *TodoUseCase) GetAllTodos() ([]Todo, error) {
	return uc.TodoRepo.IndexTodos()
}

func (uc *TodoUseCase) GetTodoByID(id int64) (Todo, error) {
	return uc.TodoRepo.ReadTodoByID(id)
}

func (uc *TodoUseCase) CreateTodo(todo Todo) (Todo, error) {
	return uc.TodoRepo.CreateTodo(todo)
}

func (uc *TodoUseCase) UpdateTodo(id int64, updatedTodo Todo) (Todo, error) {
	todo, err := uc.TodoRepo.ReadTodoByID(id)
	if err != nil {
		return Todo{}, err
	}

	if updatedTodo.Title != "" {
		todo.Title = updatedTodo.Title
	}
	if updatedTodo.Description != "" {
		todo.Description = updatedTodo.Description
	}
	if updatedTodo.Priority != 0 {
		todo.Priority = updatedTodo.Priority
	}

	return uc.TodoRepo.UpdateTodo(todo)
}

func (uc *TodoUseCase) DeleteTodo(id int64) error {
	return uc.TodoRepo.DeleteTodoByID(id)
}
