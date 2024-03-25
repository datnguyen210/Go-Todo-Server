package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var todos = MockTodos()

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", createTodo)
	router.Run("localhost:8080")
	fmt.Println("Hello users of the todo list app")
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var newTodo Todo

	// Bind JSON payload to newTodo struct
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign ID for the new to-do
	newTodo.ID = generateID()

	// Append the new to-do to the todos list
	todos = append(todos, newTodo)

	// Respond with the newly created to-do
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func generateID() string {
	// Generate a UUID version 4 (random UUID)
	id := uuid.New()

	// Convert UUID to string
	return id.String()
}
