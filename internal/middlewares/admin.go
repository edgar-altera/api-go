package middlewares

import (
	"net/http"
	"github.com/edgar-altera/api-go/internal/controllers"
	"github.com/edgar-altera/api-go/internal/models"
	"github.com/edgar-altera/api-go/pkg/helpers"
	"github.com/edgar-altera/api-go/pkg/lang"
	// log "github.com/sirupsen/logrus"
)

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := r.Context().Value("id")

		isAdmin, err := controllers.IsAdminUser(id.(int))

		if !isAdmin || err != nil {
			response := models.ErrorResponse { Success: false, Message: lang.Get("StatusForbidden", r.Header.Get("Accept-Language"))}
            
			helpers.ResponseWithJson(w, http.StatusForbidden, response)

            return
		}
		
        next.ServeHTTP(w, r)
    })
}