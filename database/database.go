package database

import (
	"datnguyen/todo/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Priority    float64   `json:"priority"`
}

func mockTodos() []Todo {
	return []Todo{
		{
			ID:          "1",
			Title:       "Complete project",
			Description: "Finish the project assignment by the end of the week.",
			Date:        time.Now().AddDate(0, 0, 7), // Due in 7 days
			Priority:    5.0,                         // High priority
		},
		{
			ID:          "2",
			Title:       "Grocery shopping",
			Description: "Buy groceries for the week.",
			Date:        time.Now().AddDate(0, 0, 2), // Due in 2 days
			Priority:    3.0,                         // Medium priority
		},
		{
			ID:          "3",
			Title:       "Exercise",
			Description: "Go for a run in the evening.",
			Date:        time.Now().AddDate(0, 0, 1), // Due tomorrow
			Priority:    2.0,                         // Low priority
		},
	}
}

var todos = mockTodos()

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")

	for _, todo := range todos {
		if todo.ID == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "No such todo"})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	for i, todo := range todos {
		if todo.ID == id {
			var updateData map[string]interface{}
			if err := c.BindJSON(&updateData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if title, ok := updateData["Title"].(string); ok {
				todos[i].Title = title
			}
			if description, ok := updateData["Description"].(string); ok {
				todos[i].Description = description
			}
			if dateStr, ok := updateData["Date"].(string); ok {
				if date, err := time.Parse("2006-01-02", dateStr); err == nil {
					todos[i].Date = date
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
			}
			if priority, ok := updateData["Priority"].(float64); ok {
				todos[i].Priority = priority
			}
			c.IndentedJSON(http.StatusOK, todos[i])

			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "No such todo"})
}

func CreateTodo(c *gin.Context) {
	var newTodo Todo

	// Bind JSON payload to newTodo struct
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign ID for the new to-do
	newTodo.ID = util.GenerateID()

	// Append the new to-do to the todos list
	todos = append(todos, newTodo)

	// Respond with the newly created to-do
	c.IndentedJSON(http.StatusCreated, newTodo)
}
