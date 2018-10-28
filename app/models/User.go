package models

import (
	"database/sql"
	"fmt"
	"strings"
)

// "fmt"
// "net/http"
// "html/template"
// "database/sql"

type User struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Type         int    `json:"type"`
	Token        string `json:"token"`
	SessionToken string `json:"session_token"`
}

// func (u User) Get() {
// 	db, _ := sql.Open("mysql", "phpmyadmin:@tcp(127.0.0.1:3306)/")

// 	defer db.Close()

// 	_,_ = db.Exec("USE questionnaire")

// 	u, err := db.Query(`
// 		SELECT * FROM users`)
// 	defer u.Close()

// 	return u
// }

func (u User) FindByFields(m map[string]string) (User, error) {
	db, err := sql.Open("mysql", "phpmyadmin:@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE questionnaire")
	if err != nil {
		panic(err)
	}

	query := "SELECT fullname, email FROM users WHERE "

	for k, v := range m { 
		query += k + "='" + v + "' AND "
	}
	query = strings.TrimSuffix(query, " AND ")

	row := db.QueryRow(query)
	err = row.Scan(&u.Fullname, &u.Email);

	return u, err
}

func (u User) Insert(user User) User {
	db, err := sql.Open("mysql", "phpmyadmin:@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE questionnaire")
	if err != nil {
		panic(err)
	}

	user.Token = u.GenerateToken()

	db.QueryRow(`
	INSERT INTO users (fullname, email, token) VALUES (?, ?, ?)`, user.Fullname, user.Email, user.Token)

	return user
}

func (u User) Update(user User) sql.Result {
	db, err := sql.Open("mysql", "phpmyadmin:@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// var db *sql.DB

	_, err = db.Exec("USE questionnaire")
	if err != nil {
		panic(err)
	}
	
	db.QueryRow(`
	UPDATE users SET username = ?, password = ?, token = NULL WHERE email = ? AND token = ?`, user.Username, user.Password, user.Email, user.Token)

	return nil
}

func (u User) GenerateToken() string {
	return "token"
}

func (u User) login() {

}

func (u User) register() {

}
