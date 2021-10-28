package main

import (
	"net/http"
	"os"
	log "github.com/sirupsen/logrus"
    "gopkg.in/natefinch/lumberjack.v2"
	"github.com/edgar-altera/api-go/internal/server"
	"github.com/joho/godotenv"
)

var port string

func main() {
	
	r := server.NewRouter()

	log.WithFields(
		log.Fields{
			"port": port,
			"bar": "bar",
		},
	).Info("Starting App")

	http.ListenAndServe(port, r)
}

func init() {

    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

	port = ":" + os.Getenv("APP_PORT")
}

func init() {
    log.SetOutput(&lumberjack.Logger{
        Filename:   os.Getenv("LOG_PATH"),
        MaxSize:    100, // megabytes
        MaxBackups: 50,
        MaxAge:     31, //days
        Compress:   false, // disabled by default
    })

    log.SetFormatter(&log.JSONFormatter{})
}