package middlewares

import (
	"net/http"
	"github.com/edgar-altera/api-go/internal/config"
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
				"AcceptLanguage": r.Header.Get("Accept-Language"),
            },
        ).Info("Logguer Request")

        next.ServeHTTP(w, r)
    })
}