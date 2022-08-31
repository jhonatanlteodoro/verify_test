package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonatanlteodoro/verify_test/handlers"
)

// Registry users routes using a prefix for every route as /users
func userRoutes(baseRouter *mux.Router) {

	r := baseRouter.PathPrefix("/users").Headers("Content-Type", "application/json").Subrouter()

	r.HandleFunc("/", handlers.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/{id}", handlers.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/", handlers.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/{id}", handlers.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/{id}", handlers.DeleteUser).Methods(http.MethodDelete)

}

func RegistryRoutes(router *mux.Router) {

	// registry user routes
	userRoutes(router)

}
