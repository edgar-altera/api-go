package main

import (
	"net/http"
	"github.com/edgar-altera/api-go/internal/config"
	"github.com/edgar-altera/api-go/internal/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	
	r := server.NewRouter()

	log.WithFields(
		log.Fields{
			"port": config.APP_PORT,
			"bar": "bar",
		},
	).Info("Starting App")

	http.ListenAndServe(config.APP_PORT, r)
}