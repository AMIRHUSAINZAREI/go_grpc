package pkg

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	envFilePath = "./.env"
)

func GetEnv(key string) (string, error) {
	// Load environment variables from .env file
	err := godotenv.Load(envFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to load environment variables from %s: %v", envFilePath, err)
	}

	// Get value of environment variable with key
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("environment variable %s not set", key)
	}

	return value, nil
}
