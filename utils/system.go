package utils

import (
	"os"
)

func GetEnv(key string, def string) string {
	/*
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	*/
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
