package controllers

import (
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/edgar-altera/api-go/internal/config"
	"github.com/edgar-altera/api-go/internal/models"
	"github.com/edgar-altera/api-go/pkg/claim"
	"github.com/edgar-altera/api-go/pkg/helpers"
)

func Login(w http.ResponseWriter, r *http.Request) {

	// To-do validate user and pass

	expiresAt := time.Now().Add(getJWTTimeExpirationUnit(config.JWT_TIME_EXPIRATION_UNIT) * time.Duration(config.JWT_TIME_EXPIRATION_VALUE)).Unix()

	c := claim.Claim{
		123, // user_id
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

    token, _ := c.NewToken(config.JWT_SEED)

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