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

	r.HandleFunc("/", controllers.AppName).Methods(http.MethodGet)

	r.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)

	sr := r.PathPrefix("/").Subrouter()

	sr.Use(middlewares.Auth)

	// sr.HandleFunc("/users", controllers.IndexUsers).Methods(http.MethodGet)

	admin := r.PathPrefix("/admin").Subrouter()

	admin.Use(middlewares.Auth)
	admin.Use(middlewares.Admin)

	admin.HandleFunc("/users", controllers.IndexUser).Methods(http.MethodGet)
	admin.HandleFunc("/users/{id}", controllers.ShowUser).Methods(http.MethodGet)
	admin.HandleFunc("/users", controllers.StoreUser).Methods(http.MethodPost)
	admin.HandleFunc("/users/{id}", controllers.UpdateUser).Methods(http.MethodPut)
	admin.HandleFunc("/users/{id}", controllers.DestroyUser).Methods(http.MethodDelete)

	return r
}