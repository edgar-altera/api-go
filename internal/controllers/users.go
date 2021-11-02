package controllers

import (
	"net/http"
	"github.com/edgar-altera/api-go/internal/config"
	"github.com/edgar-altera/api-go/internal/models"
	"github.com/edgar-altera/api-go/pkg/helpers"
	log "github.com/sirupsen/logrus"
)

 
func FindUser(w http.ResponseWriter, r *http.Request) {
	
	user := models.Person {"Edgar", 35, "Dev", true}

	response := models.Response { Success: true, Data: user}

	helpers.ResponseWithJson(w, http.StatusOK, response)

	log.WithFields(
		log.Fields{
			"name": user.Name,
			"age": user.Age,
			"degree": user.Degree,
			"test": config.APP_PORT,
		},
	).Info("Find User")

}
