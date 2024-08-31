package middlewares

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DotEnvVariable(key string) string {
	if key == "" {
		return "NULL"
	}
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
