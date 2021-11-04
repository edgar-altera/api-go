package middlewares

import (
	"net/http"
	"github.com/edgar-altera/api-go/internal/config"
    "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        id := uuid.New()

        r.Header.Set("Request-ID", id.String())

        if r.Header.Get("Accept-Language") == "" {
            r.Header.Set("Accept-Language", config.APP_LANG)
        }

        log.WithFields(
            log.Fields{
                "method": r.Method, 
                "uri": r.RequestURI,
                "host": r.Host, 
                "path": r.URL.Path,
                "remote_addr": r.RemoteAddr,
                "user_agent": r.UserAgent(),
				"accept_language": r.Header.Get("Accept-Language"),
                "request_id": r.Header.Get("Request-ID"),
            },
        ).Info("Logguer Request")

        next.ServeHTTP(w, r)
    })
}