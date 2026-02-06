package env

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		// log.Println("No .env file found or error loading .env file")
	} else {
		log.Println(".env file loaded successfully")
	}
}
