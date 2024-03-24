package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var todos = MockTodos()

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:8080")
	fmt.Println("Hello users of the todo list app")
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}
