package main;

import (
	"fmt"
	"net/http"
	"database/sql"
	"./app/controllers"
	"./app/controllers/auth"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "phpmyadmin:@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// delete later
	_,err = db.Exec(`DROP DATABASE IF EXISTS questionnaire`)
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("CREATE DATABASE IF NOT EXISTS questionnaire DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci'")
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("USE questionnaire")
	if err != nil {
		panic(err)
	}


	_,err = db.Exec(`CREATE TABLE IF NOT EXISTS users ( 
		id INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT, 
		username VARCHAR(144) NULL,
		fullname VARCHAR(144) NULL,
		email VARCHAR(144) NULL,
		password VARCHAR(255) NULL,
		type TINYINT(1) DEFAULT 1 COMMENT '0-Admin; 1-User',
		session_token VARCHAR(255) NULL
	)`)
	if err != nil {
		panic(err)
	}

	_,err = db.Exec(`CREATE TABLE IF NOT EXISTS surveys ( 
		id INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT, 
		title VARCHAR(144) NULL
	)`)
	if err != nil {
		panic(err)
	}

	_,err = db.Exec(`CREATE TABLE IF NOT EXISTS questions ( 
		id INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT, 
		title VARCHAR(144) NULL,
		content JSON NOT NULL,
		survey_id INT(10) NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	_,err = db.Exec(`CREATE TABLE IF NOT EXISTS answers ( 
		id INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT, 
		content JSON NOT NULL,
		question_id INT(10) NOT NULL,
		user_id INT(10) NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	_,err = db.Exec(`TRUNCATE TABLE users`)
	if err != nil {
		panic(err)
	}

	inser, err := db.Query(`
	INSERT INTO users (
		username, password
	) VALUES ('admin', '111111')`)
	if err != nil {
		panic(err)
	}
	defer inser.Close()

	_,err = db.Exec("USE questionnaire")
	if err != nil {
		panic(err)
	}

	insert, err := db.Query(`
	INSERT INTO users (
		username, fullname, email, password
	) VALUES ('jaskio', 'Jasmin Bektic', 'jaskio89@gmail.com', '111111')`)
	if err != nil {
		panic(err)
	}

	defer insert.Close()

	_,err = db.Exec("USE questionnaire")
	if err != nil {
		panic(err)
	}

	inserte, err := db.Query(`
	INSERT INTO questions (
		title, content, survey_id
	) VALUES ('Question 1', '{"tet": "test"}', '1')`)
	if err != nil {
		panic(err)
	}

	defer inserte.Close()
	

	
	fmt.Println("Logged in database");

	route := mux.NewRouter()

	// migrations(db);
	var l auth.LoginController
	var q controllers.QuestionController
    // l.Login()

	route.HandleFunc("/", handler);
	route.HandleFunc("/login", l.Login);
	// http.HandleFunc("/logout", l.Login);
	// http.HandleFunc("/register", l.Login);
	// http.HandleFunc("/password_reset", l.Login);

	// http.HandleFunc("/survey", q.Get).Methods("GET");
	// http.HandleFunc("/survey/{id}", q.Get).Methods("GET");
	// http.HandleFunc("/survey", q.Get).Methods("POST");
	// http.HandleFunc("/survey", q.Get).Methods("PUT");
	// http.HandleFunc("/survey/{id}", q.Get).Methods("DELETE");

	route.HandleFunc("/question", q.Get).Methods("GET");
	route.HandleFunc("/question/{id}", q.Get).Methods("GET");
	route.HandleFunc("/question", q.Insert).Methods("POST");
	route.HandleFunc("/question", q.Update).Methods("PUT");
	route.HandleFunc("/question/{id}", q.Delete).Methods("DELETE");

	// http.HandleFunc("/answer", q.Get).Methods("GET");
	// http.HandleFunc("/answer/{id}", q.Get).Methods("GET");
	// http.HandleFunc("/answer", q.Get).Methods("POST");
	// http.HandleFunc("/answer", q.Get).Methods("PUT");
	// http.HandleFunc("/answer/{id}", q.Get).Methods("DELETE");

	http.ListenAndServe(":8000", route);
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage");
}
