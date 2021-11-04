package middlewares

import (
	"fmt"
	"net/http"
	"github.com/edgar-altera/api-go/internal/config"
	"github.com/edgar-altera/api-go/internal/models"
	"github.com/edgar-altera/api-go/pkg/helpers"
	"github.com/edgar-altera/api-go/pkg/lang"
	log "github.com/sirupsen/logrus"
)

func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        if r.Header.Get("Accept-Language") == "" {
            r.Header.Set("Accept-Language", config.APP_LANG)
        }

        log.WithFields(
            log.Fields{
                "Method": r.Method, 
                "URI": r.RequestURI,
                "Host": r.Host, 
                "Path": r.URL.Path,
                "RemoteAddr": r.RemoteAddr,
                "UserAgent": r.UserAgent(),
            },
        ).Info("Logguer Request")

        next.ServeHTTP(w, r)
    })
}

func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
		authorization := r.Header.Get("Authorization")
        
        if authorization != "Bearer " + config.APP_ACCESS_TOKEN {

			response := models.ErrorResponse { Success: false, Message: lang.Get("StatusUnauthorizedMessage", r.Header.Get("Accept-Language"))}
            
			helpers.ResponseWithJson(w, http.StatusUnauthorized, response)

            return
        }

		fmt.Println("Auth middleware")

        next.ServeHTTP(w, r)
    })
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