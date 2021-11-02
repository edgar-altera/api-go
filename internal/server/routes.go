package server
 
import (
	"github.com/edgar-altera/api-go/internal/controllers"
	"github.com/edgar-altera/api-go/internal/middlewares"
	"github.com/gorilla/mux"
)
 
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.Auth)

	r.HandleFunc("/", controllers.AppName).Methods("GET")

	sr := r.PathPrefix("/").Subrouter()

	sr.Use(middlewares.User)

	sr.HandleFunc("/users", controllers.FindUser).Methods("GET")

	movies := r.PathPrefix("/movies").Subrouter()

	movies.Use(middlewares.Movie)

	movies.HandleFunc("/{id}", controllers.FindMovie).Methods("GET")
	movies.HandleFunc("/{id}", controllers.UpdateMovie).Methods("PUT")

	return r
}