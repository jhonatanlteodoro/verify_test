package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonatanlteodoro/verify_test/app/handlers"
	"gorm.io/gorm"
)

// Registry users routes using a prefix for every route as /users
func userRoutes(baseRouter *mux.Router, db *gorm.DB) {

	r := baseRouter.PathPrefix("/users").Headers("Content-Type", "application/json").Subrouter()

	r.HandleFunc("/", handlers.GetAll(db)).Methods(http.MethodGet)
	r.HandleFunc("/{id}", handlers.GetUserById(db)).Methods(http.MethodGet)
	r.HandleFunc("/", handlers.CreateUser(db)).Methods(http.MethodPost)
	r.HandleFunc("/{id}", handlers.UpdateUser(db)).Methods(http.MethodPut)
	r.HandleFunc("/{id}", handlers.DeleteUser(db)).Methods(http.MethodDelete)

}

func RegistryRoutes(router *mux.Router, db *gorm.DB) {

	// registry user routes
	userRoutes(router, db)

}
