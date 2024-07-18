// cmd/server/main.go

package main

import (
	"log"

	"datnguyen/todo/config"
	"datnguyen/todo/infra/db"
	"datnguyen/todo/internal/interface/http"
	"datnguyen/todo/internal/interface/repository"
	"datnguyen/todo/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the database connection
	database, err := db.InitMySQL(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repository
	todoRepo := repository.NewMySQLTodoRepository(database)

	// Initialize use case
	todoUseCase := usecase.NewTodoUseCase(todoRepo)

	// Initialize HTTP handler
	todoHandler := http.NewTodoHandler(todoUseCase)

	// Create a new router
	router := gin.Default()

	// Define routes
	router.GET("/todos", todoHandler.IndexTodos)
	router.GET("/todos/:id", todoHandler.ReadTodoByID)
	router.POST("/todos", todoHandler.CreateTodo)
	router.PUT("/todos/:id", todoHandler.UpdateTodo)
	router.DELETE("/todos/:id", todoHandler.DeleteTodoByID)

	// Start the server on configured port
	err = router.Run(":" + cfg.ServerPort)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
