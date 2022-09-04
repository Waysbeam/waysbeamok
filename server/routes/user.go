package routes

import (
	"Be_waysbeam/handlers"
	"Be_waysbeam/pkg/mysql"
	"Be_waysbeam/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h:= handlers.HandlersUser(userRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/user", h.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", h.DeleteUser).Methods("PATCH")
	r.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
}