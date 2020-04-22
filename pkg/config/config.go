package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *sql.DB

type mysqlconfig struct {
	host     string
	user     string
	password string
	port     string
	db       string
}

func ConnectToMysql() {
	config := mysqlconfig{
		host:     os.Getenv("Server"),
		user:     os.Getenv("UserName"),
		password: os.Getenv("Password"),
		port:     os.Getenv("Port"),
		db:       os.Getenv("DatabaseName"),
	}
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.user, config.password, config.host,
		config.port, config.db)
	d, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Error while connecting to mysql: %v", err)
		return
	}
	db = d
}


func LoadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		return
	}
	log.Println(".env file loaded successfully")
}

func GetDb() *sql.DB {
	return db
}

