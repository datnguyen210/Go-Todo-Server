package main

import (
	"fmt"

	"datnguyen/todo/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", database.GetTodos)
	router.GET("todos/:id", database.GetTodoById)
	router.POST("/todos", database.CreateTodo)
	router.PUT("/todos/:id", database.UpdateTodo)
	router.Run("localhost:8080")
	fmt.Println("Hello users of the todo list app")
}
