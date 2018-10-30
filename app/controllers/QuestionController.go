package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
)

type QuestionController struct {
}

/*
 *  Get all questions
 *  Route: /questions
 */
func (q QuestionController) GetAll(w http.ResponseWriter, r *http.Request) {
	var question models.Question
	var res []byte

	questions := question.GetAll()

	res, _ = json.Marshal(questions)
	w.Write(res)
}

/*
 *  Get question
 *  Route: /question/{id}
 */
func (q QuestionController) GetOne(w http.ResponseWriter, r *http.Request) {
	var question models.Question
	var res []byte

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	getQuestion, err := question.GetOne(id)
	if err != nil {
		res, _ = json.Marshal("Question not found.")
		w.Write(res)

		return
	}

	res, _ = json.Marshal(getQuestion)
	w.Write(res)
}

/*
 *  Create question
 *  Route: /question
 */
func (q QuestionController) Insert(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var question models.Question
	var res []byte

	err = json.Unmarshal(body, &question)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: Field validator - title, content, survey_id

	question.Insert(question)

	res, _ = json.Marshal("Question successfully created.")
	w.Write(res)
}

/*
 *  Update question
 *  Route: /question
 */
func (q QuestionController) Update(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var question models.Question
	var res []byte

	err = json.Unmarshal(body, &question)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: Field validator - id, title, content, survey_id

	question.Update(question)

	res, _ = json.Marshal("Question successfully updated.")
	w.Write(res)
}

/*
 *  Delete question
 *  Route: /question/{id}
 */
func (q QuestionController) Delete(w http.ResponseWriter, r *http.Request) {
	var question models.Question
	var res []byte

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	question.Delete(id)

	res, _ = json.Marshal("Question successfully deleted.")
	w.Write(res)
}
