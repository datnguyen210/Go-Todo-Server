package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	ServerPort string
}

// LoadConfig reads the configuration from environment variables or `.env` file.
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	config := &Config{
		DBUser:     getEnv("DBUSER", "root"),
		DBPassword: getEnv("DBPASS", ""),
		DBName:     getEnv("DBNAME", "todos"),
		DBHost:     getEnv("DBHOST", "127.0.0.1"),
		DBPort:     getEnv("DBPORT", "3306"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	return config
}

// getEnv reads an environment variable or returns a default value if not set.
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
