package usecase

import (
	"datnguyen/todo/internal/entity"
)

// TodoRepository defines the interface for interacting with the Todo data source.
type TodoRepository interface {
	IndexTodos() ([]entity.Todo, error)
	ReadTodoByID(id int64) (entity.Todo, error)
	CreateTodo(todo entity.Todo) (entity.Todo, error)
	UpdateTodo(todo entity.Todo) (entity.Todo, error)
	DeleteTodoByID(id int64) error
}

// TodoUseCase provides the application's business logic for todos.
type TodoUseCase struct {
	TodoRepo TodoRepository
}

// NewTodoUseCase creates a new instance of TodoUseCase.
func NewTodoUseCase(todoRepo TodoRepository) *TodoUseCase {
	return &TodoUseCase{TodoRepo: todoRepo}
}

// GetAllTodos retrieves all todos.
func (uc *TodoUseCase) GetAllTodos() ([]entity.Todo, error) {
	return uc.TodoRepo.IndexTodos()
}

// GetTodoByID retrieves a todo by its ID.
func (uc *TodoUseCase) GetTodoByID(id int64) (entity.Todo, error) {
	return uc.TodoRepo.ReadTodoByID(id)
}

// CreateTodo creates a new todo.
func (uc *TodoUseCase) CreateTodo(todo entity.Todo) (entity.Todo, error) {
	return uc.TodoRepo.CreateTodo(todo)
}

// UpdateTodo updates an existing todo.
func (uc *TodoUseCase) UpdateTodo(id int64, updatedTodo entity.Todo) (entity.Todo, error) {
	todo, err := uc.TodoRepo.ReadTodoByID(id)
	if err != nil {
		return entity.Todo{}, err
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

// DeleteTodo removes a todo by its ID.
func (uc *TodoUseCase) DeleteTodo(id int64) error {
	return uc.TodoRepo.DeleteTodoByID(id)
}