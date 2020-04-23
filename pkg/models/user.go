package models

import (
	"database/sql"
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
	config.ConnectToMysql()
	db = config.GetDb()
}

func (u *User) CreateUser() (int64, error) {
	q := "Insert into users(name, email, phone) values (?, ?, ?)"
	result,err := db.Exec(q, u.Name, u.Email, u.Phone)
	if err != nil {
		log.Print("Error", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		//log.Fatal("Error while inserting row: ", u) # as fatal terminates the application
		log.Print("Error while inserting row", u)
		return 0, err
	}
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
		users = append(users, user)
	}
	return users, nil
}

func UpdateUser(id int, user *User) error {
	q := "Update users set name = ? and email = ? and phone = ? where id = ?"
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