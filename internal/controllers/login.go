package controllers

import (
	"encoding/json"
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/edgar-altera/api-go/internal/config"
	"github.com/edgar-altera/api-go/internal/models"
	"github.com/edgar-altera/api-go/pkg/claim"
	"github.com/edgar-altera/api-go/pkg/helpers"
	"github.com/edgar-altera/api-go/pkg/lang"
)

func Login(w http.ResponseWriter, r *http.Request) {

	u, err := validate(r)

	if err != nil {

		response := models.ErrorResponse { Success: false, Message: lang.Get("StatusUnprocessableEntity", r.Header.Get("Accept-Language"))}
		
		helpers.ResponseWithJson(w, http.StatusUnprocessableEntity, response)

		return 
	}

	user, err := GetByUserName(u.Username)

	if err != nil {

		response := models.ErrorResponse { Success: false, Message: lang.Get("StatusNotFound", r.Header.Get("Accept-Language"))}
		
		helpers.ResponseWithJson(w, http.StatusNotFound, response)

		return 
	}

	if !user.PasswordMatch(u.Password) {

		response := models.ErrorResponse { Success: false, Message: lang.Get("StatusUnauthorized", r.Header.Get("Accept-Language"))}
		
		helpers.ResponseWithJson(w, http.StatusUnauthorized, response)

		return 
	}

	expiresAt := time.Now().Add(getJWTTimeExpirationUnit(config.JWT_TIME_EXPIRATION_UNIT) * time.Duration(config.JWT_TIME_EXPIRATION_VALUE)).Unix()

	c := claim.Claim{
		user.Id, 
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

    token, err := c.NewToken(config.JWT_SEED)

	if err != nil {

		response := models.ErrorResponse { Success: false, Message: lang.Get("StatusInternalServerError", r.Header.Get("Accept-Language"))}
		
		helpers.ResponseWithJson(w, http.StatusInternalServerError, response)

		return 
	}

	response := models.Response { Success: true, Data: token}

	helpers.ResponseWithJson(w, http.StatusOK, response)
}

func getJWTTimeExpirationUnit(unitString string) time.Duration {

	var unit time.Duration

	switch unitString {

		case "second":
			unit = time.Second
		case "minute":
			unit = time.Minute
		case "hour":
			unit = time.Hour
		default:
			unit = time.Hour
	}

	return unit
}

func validate(r *http.Request) (models.User, error) {

	var u models.User
    
	err := json.NewDecoder(r.Body).Decode(&u)

	v := validator.New()

	err = v.Struct(u)

	return u, err
}