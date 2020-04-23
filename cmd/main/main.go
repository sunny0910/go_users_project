package main

import (
	"github.com/gorilla/mux"
	"users_project/pkg/routes"

	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterUserRoutes(router)
	http.Handle("/", router)
	router.Use(headerMiddleware)
	log.Println(http.ListenAndServe(":9000", nil))
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}