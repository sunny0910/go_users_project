package main

import (
	"database/sql"
	// "encoding/json"
	// "io/ioutil"
	"net/http"
	"github.com/joho/godotenv"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"fmt"
)

var db *sql.DB
type mysqlconfig struct {
	host string
	user string
	password string
	port string
	db string
}

func loadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		return
	}
	log.Println(".env file loaded successfully")
}

func main() {
	loadEnvFile()
	config := mysqlconfig{
		host: os.Getenv("Server"),
		user: os.Getenv("UserName"),
		password: os.Getenv("Password"),
		port: os.Getenv("Port"),
		db: os.Getenv("DatabaseName"),
	}
	var err error
	db, err = connectToMysql(config)
	if err != nil {
		log.Fatalf("couldn't connect to mysql: %v", err)
		return
	}
	http.HandleFunc("/user", userHandler)
	log.Println(http.ListenAndServe(":8000", nil))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		get(w, r)
		w.WriteHeader(http.StatusOK)
	case http.MethodPost:
		post(w, r)
		w.WriteHeader(http.StatusOK)
	default :
		w.WriteHeader(http.StatusMethodNotAllowed)
	}	
}

func connectToMysql(conf mysqlconfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.user, conf.password, conf.host, conf.port, conf.db)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get"))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post"))
}