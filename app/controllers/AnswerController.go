package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../models"
)

type AnswerController struct {
}

/*
 *  Create answer
 *  Route: /answer
 */
func (a AnswerController) Insert(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var answer models.Answer
	var res []byte

	err = json.Unmarshal(body, &answer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: Field validator - content, question_id, user_id

	if answer.IsAnswered(answer) {
		res, _ := json.Marshal("You already gave answer to this question.")
		w.Write(res)

		return
	}

	answer.Insert(answer)

	res, _ = json.Marshal("Answer successfully created.")
	w.Write(res)
}
