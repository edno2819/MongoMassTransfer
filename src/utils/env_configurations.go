package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const EnvPath string = ".env"

func init() {
	loadEnvs()
}

func loadEnvs() {
	err := godotenv.Load(EnvPath)
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func GetEnvVariableDef(key string, def string) string {
	if key == "" {
		return "NULL"
	}
	err := godotenv.Load(EnvPath)

	if err != nil {
		log.Fatalf("Error loading .env file")
		return def
	}

	return os.Getenv(key)
}

func GetEnvVariable(key string) string {
	if key == "" {
		return "NULL"
	}
	err := godotenv.Load(EnvPath)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
