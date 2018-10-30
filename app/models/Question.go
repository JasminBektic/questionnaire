package models

import (
	"database/sql"
	"encoding/json"

	"../../db"
)

type Question struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Survey_id int    `json:"survey_id"`
	Content   string `json:"content"`
}

/*
 *  Get questions related to survey
 */
func (q Question) GetForSurvey(id int) []Question {
	db, err := db.Open()

	rows, err := db.Query(`SELECT id, title, content FROM questions WHERE survey_id = ?`, id)
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

/*
 *  Delete questions related to survey
 */
func (q Question) DeleteForSurvey(id int) {
	db, _ := db.Open()

	db.QueryRow(`DELETE FROM questions WHERE survey_id = ?`, id)
}

/*
 *  Get all questions
 */
func (q Question) GetAll() []Question {
	db, err := db.Open()

	rows, err := db.Query(`SELECT 
								q.id, 
								q.title, 
								q.content, 
								q.survey_id,
								s.id,
								s.title AS survey_title
							FROM 
								questions q
							LEFT JOIN
								surveys s ON q.survey_id = s.id`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var questions []Question

	for rows.Next() {
		var s Survey

		err := rows.Scan(&q.Id, &q.Title, &q.Content, &q.Survey_id, &s.Id, &s.Title)
		if err != nil {
			continue
		}

		questions = append(questions, q)
	}

	return questions
}

/*
 *  Get question
 */
func (q Question) GetOne(id int) (Question, error) {
	db, err := db.Open()

	var s Survey

	row := db.QueryRow(`SELECT 
							q.id, 
							q.title, 
							q.content, 
							q.survey_id, 
							s.id, 
							s.title AS survey_title
						FROM 
							questions q
						LEFT JOIN
							surveys s ON q.survey_id = s.id  
						WHERE q.id = ?`, id)

	err = row.Scan(&q.Id, &q.Title, &q.Content, &q.Survey_id, &s.Id, &s.Title)

	return q, err
}

/*
 *  Insert resource into questions table
 */
func (q Question) Insert(question Question) sql.Result {
	db, err := db.Open()

	insert, err := db.Prepare(`INSERT INTO 
								questions (title, content, survey_id) 
								VALUES (?, ?, ?)`)
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

/*
 *  Question update
 */
func (q Question) Update(question Question) sql.Result {
	db, err := db.Open()

	contentJSON, err := json.Marshal(question.Content)
	if err != nil {
		panic(err)
	}

	db.QueryRow(`UPDATE 
					questions 
				SET title = ?, 
					content = ?, 
					survey_id = ? 
				WHERE id = ?`, question.Title, contentJSON, question.Survey_id, question.Id)

	return nil
}

/*
 *  Delete resource from questions table
 */
func (q Question) Delete(id int) sql.Result {
	db, _ := db.Open()

	db.QueryRow(`DELETE FROM questions WHERE id = ?`, id)

	return nil
}
