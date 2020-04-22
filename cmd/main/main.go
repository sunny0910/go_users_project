package main

import (
	"github.com/gorilla/mux"
	"users_project/pkg/config"
	"users_project/pkg/routes"

	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.LoadEnvFile()
	config.ConnectToMysql()
	router := mux.NewRouter()
	routes.RegisterUserRoutes(router)
	log.Println(http.ListenAndServe(":9000", nil))
}

