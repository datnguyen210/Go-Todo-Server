package util

import "github.com/google/uuid"

func GenerateID() string {
	// Generate a UUID version 4 (random UUID)
	id := uuid.New()

	// Convert UUID to string
	return id.String()
}
