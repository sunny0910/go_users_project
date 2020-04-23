package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"users_project/pkg/models"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user = models.GetUser(vars["id"])
	resp, _ := json.Marshal(user)
	_, _ = w.Write(resp)
}

func GetUsers(w http.ResponseWriter, r *http.Request)  {
	users, err := models.GetAllUsers()
	if err != nil {
		log.Println("error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("something unexpected happen"))
		return
	}
	resp, _ := json.Marshal(users)
	_, _ = w.Write(resp)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &user)
	id, err := user.CreateUser()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	_, _ = w.Write([]byte(fmt.Sprintf("User created with id %v", id)))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		_, _ = w.Write([]byte("Error parsing id"))
		return
	}
	_ = json.Unmarshal(body, &user)
	er := models.UpdateUser(userId, &user)
	if er != nil {
		_, _ = w.Write([]byte("Error while updating"))
		return
	}
	user.Id = userId
	resp, _ := json.Marshal(user)
	_, _ = w.Write(resp)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		_, _ = w.Write([]byte("Error parsing id"))
		return
	}
	er := models.DeleteUser(userId)
	if er != nil {
		_, _ = w.Write([]byte("Error while deleting user"))
		return
	}
	_, _ = w.Write([]byte("User deleted"))
}

