package models

import (  
	// "fmt"
	// "net/http"
	// "html/template"
	// "database/sql"
)

type User struct {  
    FirstName   string
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