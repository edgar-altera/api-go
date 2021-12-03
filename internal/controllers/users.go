package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/edgar-altera/api-go/internal/config"
	"github.com/edgar-altera/api-go/internal/database"
	"github.com/edgar-altera/api-go/internal/models"
	"github.com/edgar-altera/api-go/pkg/helpers"
	"github.com/edgar-altera/api-go/pkg/lang"
	log "github.com/sirupsen/logrus"
)

func IndexUser(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	id := r.Context().Value("id")
	fmt.Printf("ID from context %d \n", id)
	users, err := index()
	
	if err != nil {

		response := models.ErrorResponse { Success: false, Message: lang.Get("StatusInternalServerError", r.Header.Get("Accept-Language"))}
		
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

func ShowUser(w http.ResponseWriter, r *http.Request) {

}

func StoreUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DestroyUser(w http.ResponseWriter, r *http.Request) {

}

func index() ([]models.User, error) {

	db := database.DB

	var users [] models.User

    var user models.User

	rows, err := db.Query(fmt.Sprintf("SELECT id, username, status, config, created_at, updated_at FROM %s", models.UsersTable))
 
	if err != nil {
			
		return nil, err
	}

	for rows.Next() {
		
		err = rows.Scan(&user.Id, &user.Username, &user.Status, &user.Config, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			
			return nil, err
		}

		users = append(users, user)
	}

    rows.Close()

	return users, nil
}

func GetByUserName(username string) (models.User, error) {

	db := database.DB

    var user models.User

	row, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE username='%s' ", models.UsersTable, username))
 
	if err != nil {
			
		return user, err
	}

	row.Next()

	err = row.Scan(&user.Id, &user.Username, &user.PasswordHash, &user.Status, &user.IsAdmin, &user.Config, &user.CreatedAt, &user.UpdatedAt)

	row.Close()

	return user, err
}

func IsAdminUser(id int) (bool, error) {

	db := database.DB

    var count int

	err := db.QueryRow(fmt.Sprintf("SELECT count(*) FROM %s WHERE id=%d AND is_admin=true ", models.UsersTable, id)).Scan(&count)
 
	if err != nil {
			
		return false, err
	}

	if count == 1 {

		return true, nil
	}
	
	return false, errors.New("User is not Admin")
}

