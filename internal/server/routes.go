package server
 
import (
	"net/http"
	"github.com/edgar-altera/api-go/internal/controllers"
	"github.com/edgar-altera/api-go/internal/middlewares"
	"github.com/gorilla/mux"
)
 
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.Logger)

	r.Use(middlewares.Auth)

	r.HandleFunc("/", controllers.AppName).Methods(http.MethodGet)

	sr := r.PathPrefix("/").Subrouter()

	sr.Use(middlewares.User)

	sr.HandleFunc("/users", controllers.AllUsers).Methods(http.MethodGet)

	movies := r.PathPrefix("/movies").Subrouter()

	movies.Use(middlewares.Movie)

	movies.HandleFunc("/{id}", controllers.FindMovie).Methods(http.MethodGet)
	movies.HandleFunc("/{id}", controllers.UpdateMovie).Methods(http.MethodPut)

	return r
}