package load

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {

	err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func Start() {
	
	log.Println("Starting godotenv")
}