package controllers

import (
	"fmt"
	"net/http"
	"github.com/edgar-altera/api-go/pkg/helpers"
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
	// fmt.Fprintln(w, "Delete movie!!")
	helpers.ResponseWithJson(w, http.StatusOK, map[string]string{"result": "success", "id": "a45f4osaask2"})
}
