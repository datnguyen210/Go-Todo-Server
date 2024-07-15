package main

import (
	"datnguyen/todo/database"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func stringToInt(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting string to int64: %w", err)
	}
	return i, nil
}

func main() {
	// Initialize the database connection
	database.DatabaseInit()
	println("Database Initialized!")

	// Create a new router
	router := gin.Default()

	router.GET("/todos", indexTodos)
	router.GET("/todos/:id", readTodo)

	// Start the server on localhost:8080
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func indexTodos(c *gin.Context) {
	todos, err := database.IndexTodos()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func readTodo(c *gin.Context) {
	id, _ := stringToInt(c.Param("id"))

	todo, err := database.ReadTodo(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}
