package models

import (
	"database/sql"
	"encoding/json"
)

type Answer struct {
	Id          int    `json:"id"`
	Content     string `json:"content"`
	Question_id int    `json:"question_id"`
	User_id     int    `json:"user_id"`
}

/*
 *  Insert resource into answers table
 */
func (a Answer) Insert(answer Answer) sql.Result {
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

	insert, err := db.Prepare(`
	INSERT INTO answers (content, question_id, user_id) VALUES (?, ?, ?)`)
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	contentJSON, err := json.Marshal(answer.Content)
	if err != nil {
		panic(err)
	}

	res, err := insert.Exec(contentJSON, answer.Question_id, answer.User_id)
	if err != nil {
		panic(err)
	}

	return res
}

/*
 *  Get one survey with appropriate questions
 */
func (a Answer) IsAnswered(answer Answer) bool {
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

	row := db.QueryRow(`SELECT id FROM answers WHERE question_id = ? AND user_id = ?`, answer.Question_id, answer.User_id)
	err = row.Scan(&a.Id)
	if err != nil {
		return false
	}

	return true
}
