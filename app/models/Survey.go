package models

import (
	"../../db"
)

type Survey struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

/*
 *  Get all survey with appropriate questions
 */
func (s Survey) GetAll() []Survey {
	db, err := db.Open()

	rows, err := db.Query(`SELECT
							  id, 
							  title
						  FROM 
							  surveys`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var surveys []Survey

	for rows.Next() {
		var s Survey
		var q Question

		err := rows.Scan(&s.Id, &s.Title)
		if err != nil {
			continue
		}

		s.Questions = q.GetForSurvey(s.Id)

		surveys = append(surveys, s)
	}

	return surveys
}

/*
 *  Get one survey with appropriate questions
 */
func (s Survey) GetOne(id int) (Survey, error) {
	db, err := db.Open()

	var q Question

	row := db.QueryRow(`SELECT id, title FROM surveys WHERE id = ?`, id)
	err = row.Scan(&s.Id, &s.Title)

	s.Questions = q.GetForSurvey(s.Id)

	return s, err
}

/*
 *  Insert resource into surveys table
 */
func (s Survey) Insert(survey Survey) {
	db, _ := db.Open()

	db.QueryRow(`INSERT INTO surveys (title) VALUES (?)`, survey.Title)
}

/*
 *  Survey update
 */
func (s Survey) Update(survey Survey) {
	db, _ := db.Open()

	db.QueryRow(`UPDATE surveys SET title = ? WHERE id = ?`, survey.Title, survey.Id)
}

/*
 *  Delete resource from surveys table
 */
func (s Survey) Delete(id int) {
	db, _ := db.Open()

	db.QueryRow(`DELETE FROM surveys WHERE id = ?`, id)
}
