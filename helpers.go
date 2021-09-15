package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVal(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
