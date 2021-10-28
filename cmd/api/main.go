package main

import (
	"net/http"
	log "github.com/sirupsen/logrus"
    "gopkg.in/natefinch/lumberjack.v2"
	"github.com/edgar-altera/api-go/internal/load"
	"github.com/edgar-altera/api-go/internal/config"
	"github.com/edgar-altera/api-go/internal/server"
)

var port string

func main() {
	
	r := server.NewRouter()

	log.WithFields(
		log.Fields{
			"port": port,
			"bar": "bar",
			"test": config.APP_PORT,
		},
	).Info("Starting App")

	http.ListenAndServe(port, r)
}

func init() {

	load.Start()

	port = ":" + config.APP_PORT
}

func init() {
    log.SetOutput(&lumberjack.Logger{
        Filename:   config.LOG_PATH,
        MaxSize:    100, // megabytes
        MaxBackups: 50,
        MaxAge:     31, //days
        Compress:   false, // disabled by default
    })

    log.SetFormatter(&log.JSONFormatter{})
}