package utils

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	TURSO_DATABASE_URL string
	TURSO_AUTH_TOKEN   string
	JWT_SECRET         string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}

	TURSO_DATABASE_URL = os.Getenv("TURSO_DATABASE_URL")
	TURSO_AUTH_TOKEN = os.Getenv("TURSO_AUTH_TOKEN")
	JWT_SECRET = os.Getenv("JWT_SECRET")
}



