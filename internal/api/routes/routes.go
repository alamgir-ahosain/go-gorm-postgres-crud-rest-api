package routes

import (
	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/api/handlers"
	getusers "github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/api/handlers/get_users"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", getusers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getusers.GetUsersByID).Methods("GET")

	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")

	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")

	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
}
