package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ConnectionString() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv("CONN_STR")
}
