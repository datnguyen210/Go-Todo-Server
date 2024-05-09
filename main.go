package main

import (
	"datnguyen/todo/database"
)

func main() {
	// Initialize the database connection
	database.DatabaseInit()

	println("Initialized!")
}
