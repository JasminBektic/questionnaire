package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
)

type SurveyController struct {
}

/*
 *  Get all surveys
 *  Route: /survey
 */
func (s SurveyController) GetAll(w http.ResponseWriter, r *http.Request) {
	var survey models.Survey
	var res []byte

	surveys := survey.GetAll()

	res, _ = json.Marshal(surveys)
	w.Write(res)
}

/*
 *  Get survey
 *  Route: /survey/{id}
 */
func (s SurveyController) GetOne(w http.ResponseWriter, r *http.Request) {
	var survey models.Survey
	var res []byte

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	getSurvey, err := survey.GetOne(id)
	if err != nil {
		res, _ = json.Marshal("Survey not found.")
		w.Write(res)

		return
	}

	res, _ = json.Marshal(getSurvey)
	w.Write(res)
}

/*
 *  Create survey
 *  Route: /survey
 */
func (s SurveyController) Insert(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var survey models.Survey
	var res []byte

	err = json.Unmarshal(body, &survey)
	if err != nil {
		return
	}

	// TODO: Field validator - title

	survey.Insert(survey)

	res, _ = json.Marshal("Survey successfully created.")
	w.Write(res)
}

/*
 *  Update survey
 *  Route: /survey
 */
func (s SurveyController) Update(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var survey models.Survey
	var res []byte

	err = json.Unmarshal(body, &survey)
	if err != nil {
		return
	}

	// TODO: Field validator - id, title

	survey.Update(survey)

	res, _ = json.Marshal("Survey successfully updated.")
	w.Write(res)
}

/*
 *  Delete survey
 *  Route: /delete/{id}
 */
func (s SurveyController) Delete(w http.ResponseWriter, r *http.Request) {
	var survey models.Survey
	var question models.Question
	var res []byte

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	survey.Delete(id)
	question.DeleteForSurvey(id)

	res, _ = json.Marshal("Survey successfully deleted.")
	w.Write(res)
}
