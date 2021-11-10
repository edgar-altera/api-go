package controllers

import (
	"fmt"
	"net/http"
	"github.com/edgar-altera/api-go/internal/config"
	"github.com/edgar-altera/api-go/internal/database"
	"github.com/edgar-altera/api-go/internal/models"
	"github.com/edgar-altera/api-go/pkg/helpers"
	"github.com/edgar-altera/api-go/pkg/lang"
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

func AllUsers(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	id := r.Context().Value("id")
	fmt.Printf("ID from context %d \n", id)
	users, err := index()
	
	if err != nil {

		response := models.ErrorResponse { Success: false, Message: lang.Get("StatusInternalServerErrorMessage", r.Header.Get("Accept-Language"))}
		
		helpers.ResponseWithJson(w, http.StatusInternalServerError, response)

		log.WithFields(
			log.Fields{
				"error": err,
				"request_id": r.Header.Get("Request-ID"),
			},
		).Error("AllUsers")

		return
	}

	response := models.Response { Success: true, Data: users}

	helpers.ResponseWithJson(w, http.StatusOK, response)

	log.WithFields(
		log.Fields{
			"test": config.APP_PORT,
			"request_id": r.Header.Get("Request-ID"),
		},
	).Info("AllUsers")

}

func index() ([]models.User, error) {

	db := database.DB

	users := [] models.User {}

    var user models.User

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", models.UsersTable))
 
	if err != nil {
			
		return nil, err
	}

	for rows.Next() {
		
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Status, &user.Config, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			
			return nil, err
		}

		users = append(users, user)
	}

    rows.Close()

	return users, nil
}
