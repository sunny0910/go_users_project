package routes

import (
	"users_project/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
}
