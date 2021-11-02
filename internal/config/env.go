package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"fmt"
)

var APP_NAME string
var APP_PORT string
var APP_ACCESS_TOKEN string
var LOG_PATH string

func init() {

	err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

	fmt.Println("Load env.go")
}

func init() {

	APP_NAME = os.Getenv("APP_NAME")
	APP_PORT = os.Getenv("APP_PORT")
	APP_ACCESS_TOKEN = os.Getenv("APP_ACCESS_TOKEN")
	LOG_PATH = os.Getenv("LOG_PATH")

	fmt.Printf("Load vars log_path=%s \n", LOG_PATH)
	
}