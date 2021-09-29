package main

import (
	"net/http"
	"log"
	"os"
	"github.com/edgar-altera/api-go/internal/server"
	"github.com/joho/godotenv"
)

var port string

func main() {
	
	r := server.NewRouter()
 
	log.Printf("Starting App on port %v", port)

	http.ListenAndServe(port, r)
}

func init() {

    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

	port = ":" + os.Getenv("APP_PORT")
}