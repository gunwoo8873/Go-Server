package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Private ENV configure file
func ReadEnv() {
	// Current work get directory path
	workPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// ENV file path is get from work path directory
	envPath := filepath.Join(workPath, ".env")

	// Read to string env data
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}

	return value
}
