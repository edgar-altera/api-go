package middlewares

import (
    "context"
    "errors"
    "strings"
	"fmt"
	"net/http"
	"github.com/edgar-altera/api-go/internal/config"
	"github.com/edgar-altera/api-go/internal/models"
	"github.com/edgar-altera/api-go/pkg/claim"
	"github.com/edgar-altera/api-go/pkg/helpers"
	"github.com/edgar-altera/api-go/pkg/lang"
	// log "github.com/sirupsen/logrus"
)

func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
		authorization := r.Header.Get("Authorization")

        tokenString, err := tokenFromAuthorization(authorization)

        if err != nil {

            response := models.ErrorResponse { Success: false, Message: lang.Get("StatusUnauthorizedMessage", r.Header.Get("Accept-Language"))}
            
			helpers.ResponseWithJson(w, http.StatusUnauthorized, response)

            return
        }

        signingString := config.JWT_SEED

        c, err := claim.GetFromToken(tokenString, signingString)

        if err != nil {
            
            response := models.ErrorResponse { Success: false, Message: lang.Get("StatusUnauthorizedMessage", r.Header.Get("Accept-Language"))}
            
			helpers.ResponseWithJson(w, http.StatusUnauthorized, response)

            return
        }

        ctx := r.Context()
        ctx = context.WithValue(ctx, "id", c.ID)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func tokenFromAuthorization(authorization string) (string, error) {

    if authorization == "" {
        
        return "", errors.New("Autorization is required")
    }

    if !strings.HasPrefix(authorization, "Bearer") {
        
        return "", errors.New("Invalid autorization format")
    }

    l := strings.Split(authorization, " ")
    
    if len(l) != 2 {
        
        return "", errors.New("Invalid autorization format")
    }

    return l[1], nil
}

func User(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // fmt.Println("middleware", r.URL)
		fmt.Println("Test middleware")
        h.ServeHTTP(w, r)
    })
}

func Movie(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // fmt.Println("middleware", r.URL)
		fmt.Println("Test middleware 2")
        h.ServeHTTP(w, r)
    })
}