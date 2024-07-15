package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Todo struct {
	ID          int64
	Title       string
	Description string
	Priority    float64
}

func DatabaseInit() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "todos",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func IndexTodos() ([]Todo, error) {
	var todos []Todo

	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, fmt.Errorf("GetAllTodos%v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var todo Todo
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

func ReadTodo(id int64) (Todo, error) {
	var todo Todo

	row := db.QueryRow("SELECT * FROM todo WHERE id = ?", id)
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Priority); err != nil {
		if err == sql.ErrNoRows {
			return todo, fmt.Errorf("no data found with id %d", id)
		}
		return todo, fmt.Errorf("TodoById %d : %v", id, err)
	}
	return todo, nil
}

func UpdateTodo(id int64, newTodo Todo) (Todo, error) {
	existingTodo, err := ReadTodo(id)
	if err != nil {
		return existingTodo, err
	}
	if newTodo.Title != "" {
		existingTodo.Title = newTodo.Title
	}
	if newTodo.Description != "" {
		existingTodo.Description = newTodo.Description
	}
	if newTodo.Priority != 0 {
		existingTodo.Priority = newTodo.Priority
	}

	// Update the to do in the database
	_, err = db.Exec("UPDATE todo SET title = ?, description = ?, priority = ? WHERE id = ?", existingTodo.Title, existingTodo.Description, existingTodo.Priority, existingTodo.ID)
	if err != nil {
		return existingTodo, fmt.Errorf("UpdateTodo: %v", err)
	}

	return existingTodo, nil
}

func CreateTodo(todo Todo) (Todo, error) {
	createdTodo, err := db.Exec("INSERT INTO todo (title, description, priority) VALUES (?, ?, ?)", todo.Title, todo.Description, todo.Priority)
	if err != nil {
		return Todo{}, fmt.Errorf("CreateTodo: %v", err)
	}

	id, err := createdTodo.LastInsertId()
	if err != nil {
		return Todo{}, fmt.Errorf("CreateTodo: %v", err)
	}

	todo.ID = id
	return todo, nil
}

func DeleteTodo(id int64) error {
	_, err := db.Exec("DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("DeleteTodo: %v", err)
	}

	return nil
}
