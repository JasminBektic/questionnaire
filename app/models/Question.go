package models

import (
	// "fmt"
	// "net/http"
	"database/sql"
	"encoding/json"
)

type Question struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Survey_id int    `json:"survey_id"`
	Content   string `json:"content"`
}

func (q Question) GetForSurvey(id int) []Question {
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

	// var rows Question
	rows, err := db.Query(`
	SELECT id, title, content FROM questions WHERE survey_id = ?`, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var questions []Question

	for rows.Next() {
		err := rows.Scan(&q.Id, &q.Title, &q.Content)
		if err != nil {
			continue
		}

		questions = append(questions, q)
	}

	return questions
}

func (q Question) DeleteForSurvey(id int) {
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
	DELETE FROM questions WHERE survey_id = ?`, id)
}

func (q Question) Get() []byte {
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

	// var rows Question
	rows, err := db.Query(`
	SELECT title, content, survey_id FROM questions`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var questions []Question

	for rows.Next() {
		err := rows.Scan(&q.Title, &q.Content, &q.Survey_id)
		if err != nil {
			continue
		}

		questions = append(questions, q)
	}

	qs, _ := json.Marshal(&questions)

	return qs
}

func (q Question) Insert(question Question) sql.Result {
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

	// var rows Question
	insert, err := db.Prepare(`
	INSERT INTO questions (title, content, survey_id) VALUES (?, ?, ?)`)
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	contentJSON, err := json.Marshal(question.Content)
	if err != nil {
		panic(err)
	}

	res, err := insert.Exec(question.Title, contentJSON, question.Survey_id)
	if err != nil {
		panic(err)
	}

	return res
}

func (q Question) Update(question Question) sql.Result {
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

	contentJSON, err := json.Marshal(question.Content)
	if err != nil {
		panic(err)
	}

	// var rows Question
	db.QueryRow(`
	UPDATE questions SET title = ?, content = ? WHERE id = ?`, question.Title, contentJSON, question.Id)

	return nil
}

func (q Question) Delete(id int) sql.Result {
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
	DELETE FROM questions WHERE id = ?`, id)

	return nil
}
