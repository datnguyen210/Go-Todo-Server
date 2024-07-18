package repository

import (
	"database/sql"
	"fmt"

	"datnguyen/todo/internal/entity"
)

type MySQLTodoRepository struct {
	DB *sql.DB
}

func NewMySQLTodoRepository(db *sql.DB) *MySQLTodoRepository {
	return &MySQLTodoRepository{DB: db}
}

func (r *MySQLTodoRepository) IndexTodos() ([]entity.Todo, error) {
	var todos []entity.Todo

	rows, err := r.DB.Query("SELECT id, title, description, priority FROM todo")
	if err != nil {
		return nil, fmt.Errorf("GetAllTodos%v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var todo entity.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Priority); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}
	return todos, nil
}

func (r *MySQLTodoRepository) ReadTodoByID(id int64) (entity.Todo, error) {
	var todo entity.Todo

	row := r.DB.QueryRow("SELECT id, title, description, priority FROM todo WHERE id = ?", id)
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Priority); err != nil {
		if err == sql.ErrNoRows {
			return todo, fmt.Errorf("no data found with id %d", id)
		}
		return todo, fmt.Errorf("TodoById %d : %v", id, err)
	}
	return todo, nil
}

func (r *MySQLTodoRepository) CreateTodo(todo entity.Todo) (entity.Todo, error) {
	result, err := r.DB.Exec("INSERT INTO todo (title, description, priority) VALUES (?, ?, ?)", todo.Title, todo.Description, todo.Priority)
	if err != nil {
		return entity.Todo{}, fmt.Errorf("CreateTodo: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.Todo{}, fmt.Errorf("CreateTodo: %v", err)
	}

	todo.ID = id
	return todo, nil
}

func (r *MySQLTodoRepository) UpdateTodo(todo entity.Todo) (entity.Todo, error) {
	_, err := r.DB.Exec("UPDATE todo SET title = ?, description = ?, priority = ? WHERE id = ?", todo.Title, todo.Description, todo.Priority, todo.ID)
	if err != nil {
		return todo, fmt.Errorf("UpdateTodo: %v", err)
	}

	return todo, nil
}

func (r *MySQLTodoRepository) DeleteTodoByID(id int64) error {
	_, err := r.DB.Exec("DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("DeleteTodo: %v", err)
	}

	return nil
}
