package main

import (
    "time"
)

// Todo represents a todo item in the todo list.
type Todo struct {
    ID          string    `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Date        time.Time `json:"date"`
    Priority    float64   `json:"priority"`
}

// MockTodos returns mock data for three todos.
func MockTodos() []Todo {
    return []Todo{
        {
            ID:          "1",
            Title:       "Complete project",
            Description: "Finish the project assignment by the end of the week.",
            Date:        time.Now().AddDate(0, 0, 7), // Due in 7 days
            Priority:    5.0,                          // High priority
        },
        {
            ID:          "2",
            Title:       "Grocery shopping",
            Description: "Buy groceries for the week.",
            Date:        time.Now().AddDate(0, 0, 2), // Due in 2 days
            Priority:    3.0,                          // Medium priority
        },
        {
            ID:          "3",
            Title:       "Exercise",
            Description: "Go for a run in the evening.",
            Date:        time.Now().AddDate(0, 0, 1), // Due tomorrow
            Priority:    2.0,                          // Low priority
        },
    }
}
