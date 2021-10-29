package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"fmt"
)

var APP_PORT string
var LOG_PATH string

func init() {

	err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

	fmt.Println("Load env.go")
}

func init() {

	APP_PORT = os.Getenv("APP_PORT")
	LOG_PATH = os.Getenv("LOG_PATH")

	fmt.Printf("Load vars log_path=%s \n", LOG_PATH)
	
}