package controllers

import (
	"fmt"
	"net/http"
	"github.com/edgar-altera/api-go/internal/models"
	"github.com/edgar-altera/api-go/pkg/helpers"
	log "github.com/sirupsen/logrus"
)

func AllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented !")
}
 
func FindMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented !")
}
 
func CreateMovie(w http.ResponseWriter, t *http.Request) {
	fmt.Fprintln(w, "not implemented !")
}
 
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented !")
}
 
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	
	// p := models.Person {"Edgar", 35, "Dev", true}

	a := models.Address {"Armando Carrera", 5218, false}

	response := models.Response { Success: true, Data: a}

	helpers.ResponseWithJson(w, http.StatusOK, response)

	log.WithFields(
		log.Fields{
			"address": a.Street,
			"number": a.Number,
			"status": a.Status,
		},
	).Info("Delete movie")

}
