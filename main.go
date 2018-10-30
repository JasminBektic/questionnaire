package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"./app/controllers"
	"./app/controllers/auth"
	"./app/middleware"

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
	_, err = db.Exec(`DROP DATABASE IF EXISTS questionnaire`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS questionnaire DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci'")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE questionnaire")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users ( 
		id INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT, 
		username VARCHAR(50) NULL,
		fullname VARCHAR(50) NULL,
		email VARCHAR(80) NULL,
		password VARCHAR(255) NULL,
		type TINYINT(1) DEFAULT 1 COMMENT '0-Admin; 1-User',
		token VARCHAR(255) NULL,
		session_token VARCHAR(255) NULL
	)`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS surveys ( 
		id INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT, 
		title VARCHAR(144) NULL
	)`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS questions ( 
		id INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT, 
		title VARCHAR(144) NULL,
		content JSON NOT NULL,
		survey_id INT(10) NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS answers ( 
		id INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT, 
		content JSON NOT NULL,
		question_id INT(10) NOT NULL,
		user_id INT(10) NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS password_resets ( 
		email VARCHAR(144) NOT NULL,
		token VARCHAR(255) NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`TRUNCATE TABLE users`)
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

	_, err = db.Exec("USE questionnaire")
	if err != nil {
		panic(err)
	}

	insert, err := db.Query(`
	INSERT INTO users (
		username, fullname, email, password, session_token
	) VALUES ('jaskio', 'Jasmin Bektic', 'jaskio89@gmail.com', '$2a$10$O.WSJjmRfwuwwxPxSKQPaOEvnZPE0Pi8i/MvZdEb4TBPdzzdDuidi', 'token')`)
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	_, err = db.Exec("USE questionnaire")
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

	_, err = db.Exec("USE questionnaire")
	if err != nil {
		panic(err)
	}

	ins, err := db.Query(`
	INSERT INTO surveys (
		title
	) VALUES ('Survey 1')`)
	if err != nil {
		panic(err)
	}
	defer ins.Close()

	fmt.Println("Logged in database")

	router := mux.NewRouter()

	// migrations(db);
	var l auth.LoginController
	var r auth.RegisterController
	var pr auth.PasswordResetController
	var q controllers.QuestionController
	var s controllers.SurveyController
	var ts middleware.TokenSessionMiddleware
	var al middleware.AuthLevelMiddleware

	// Middleware
	router.Use(ts.Handle)
	router.Use(al.Handle)

	// Routes
	router.HandleFunc("/login", l.Login)
	router.HandleFunc("/logout", l.Logout)
	router.HandleFunc("/register", r.Register).Methods("POST")
	router.HandleFunc("/register/finish/{email}/{token}", r.FinishRegistration).Methods("POST")
	// router.HandleFunc("/password/reset", pr.ResetRequest).Methods("GET")
	router.HandleFunc("/password/reset", pr.ResetRequest).Methods("POST")
	// router.HandleFunc("/password/reset/{email}/{token}", pr.ResetFinish).Methods("GET")
	router.HandleFunc("/password/reset/{email}/{token}", pr.ResetFinish).Methods("POST")

	router.HandleFunc("/survey", s.GetAll).Methods("GET")
	router.HandleFunc("/survey/{id}", s.GetOne).Methods("GET")
	router.HandleFunc("/survey", s.Insert).Methods("POST")
	router.HandleFunc("/survey", s.Update).Methods("PUT")
	router.HandleFunc("/survey/{id}", s.Delete).Methods("DELETE")

	router.HandleFunc("/question", q.GetAll).Methods("GET")
	router.HandleFunc("/question/{id}", q.GetOne).Methods("GET")
	router.HandleFunc("/question", q.Insert).Methods("POST")
	router.HandleFunc("/question", q.Update).Methods("PUT")
	router.HandleFunc("/question/{id}", q.Delete).Methods("DELETE")

	// http.HandleFunc("/answer", q.Get).Methods("GET");
	// http.HandleFunc("/answer/{id}", q.Get).Methods("GET");
	// http.HandleFunc("/answer", q.Get).Methods("POST");
	// http.HandleFunc("/answer", q.Get).Methods("PUT");
	// http.HandleFunc("/answer/{id}", q.Get).Methods("DELETE");

	http.ListenAndServe(":8000", router)
}
