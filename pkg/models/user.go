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

func GetUser(id string) (User, error) {
	row := db.QueryRow("select * from users u where u.id = ?", id)
	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Printf("No users found for id: %v", id)
		return user, err
	}
	return user, nil
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

func UpdateUser(id int, user *User) (int64, error) {
	q := "Update users set name = ?, email = ?, phone = ? where id = ?"
	result, err := db.Exec(q, user.Name, user.Email, user.Phone, id)
	rows, _ := result.RowsAffected()
	if err != nil || rows == 0 {
		log.Printf("Error while updating user: %v or User not found id: %v | error: %v", user, id, err)
		return 0, err
	}
	return rows, nil
}

func DeleteUser(id int) (int64, error) {
	q := "Delete from users where id = ?"
	result, err := db.Exec(q, id)
	deleted, _ := result.RowsAffected()
	if err != nil || deleted == 0 {
		log.Printf("Error while deleting user : %v", id)
		return 0, err
	}
	return deleted, nil
}