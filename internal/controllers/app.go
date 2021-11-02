package controllers

import (
	"net/http"
	"github.com/edgar-altera/api-go/internal/config"
	"github.com/edgar-altera/api-go/internal/models"
	"github.com/edgar-altera/api-go/pkg/helpers"
	log "github.com/sirupsen/logrus"
)

func AppName(w http.ResponseWriter, r *http.Request) {

	response := models.Response { Success: true, Data: config.APP_NAME}

	helpers.ResponseWithJson(w, http.StatusOK, response)

	log.WithFields(
		log.Fields{
			"app": config.APP_NAME,
			"port": config.APP_PORT,
		},
	).Info("App Index Path")

}
