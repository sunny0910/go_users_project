package models

import (
	"database/sql"
	"fmt"
	"log"
	"users_project/pkg/config"
)

type User struct {
	Id        int
	Name      string
	Email     string
	Phone     string
	CreatedAt string
	UpdatedAt string
}

var db *sql.DB

func init() {
	db = config.GetDb()
}

func (u *User) CreateUser() (int, error) {
	q := "Insert into users(name, email, phone) values (?, ?, ?); Select LAST_INSERT_ID();"
	row := db.QueryRow(q, u.Name, u.Email, u.Phone)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Fatalf("Error while inserting row: %v, %v, %v", u.Name,
		u.Email, u.Phone)
		return 0, err
	}
	fmt.Print(id)
	return id, nil
}

func GetUser(id string) User {
	row := db.QueryRow("select * from users u where u.id = ?", id)
	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Fatalf("No users found for id: %v", id)
	}
	return user
}

func GetAllUsers() ([]User, error) {
	rows, err := db.Query("select * from users")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
		fmt.Println(err)
		users = append(users, user)
	}
	return users, nil
}

func UpdateUser(id int, user *User) error {
	q := "Update users set name = ? and email = ? and phone =? where id = ?"
	_, err := db.Exec(q, user.Name, user.Email, user.Phone, id)
	if err != nil {
		log.Fatal("Error while updating user: ", user)
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	q := "Delete from users where id = ?"
	_, err := db.Exec(q, id)
	if err != nil {
		log.Fatalf("Error while deleting user : %v", id)
		return err
	}
	return nil
}