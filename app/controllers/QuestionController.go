package controllers

import (  
	"fmt"
	"net/http"
	"io/ioutil"
	// "html/template"
	// "database/sql"
	"encoding/json"
	"../models"
	"strconv"

	"github.com/gorilla/mux"
)

type QuestionController struct {  
	// questionModel Question
	Questions models.Question
}

// func New(questionModel Question) questionModel {  
//     e := questionModel {firstName, lastName, totalLeave, leavesTaken}
//     return e
// }

func (q QuestionController) Get(w http.ResponseWriter, r *http.Request) {
	var question models.Question
	questions := question.Get()

	// t,_ := template.ParseFiles("templates/question.html")
	// t := template.New("fieldname example")
    // t, _ = t.Parse("hello {{.Questions}}!")

	// w.Header().Set("Content-Type", "text/html")
	w.Write(questions)
	
	// p := QuestionController{Questions: questions}
	// t.Execute(w, p)
}

func (q QuestionController) Insert(w http.ResponseWriter, r *http.Request) {
	// test, _ := json.Marshal(&r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// api.RequestFailed(h.res, w, err)
		return
	}

	var question models.Question

	err = json.Unmarshal(body, &question)
	if err != nil {
		// api.ParseFailed(h.res, w, err)
		return
	}
	insert := question.Insert(question)
	fmt.Println(insert)
}

func (q QuestionController) Update(w http.ResponseWriter, r *http.Request) {
	// test, _ := json.Marshal(&r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// api.RequestFailed(h.res, w, err)
		return
	}

	var question models.Question

	err = json.Unmarshal(body, &question)
	if err != nil {
		// api.ParseFailed(h.res, w, err)
		return
	}
	update := question.Update(question)
	fmt.Println(update)
}

func (q QuestionController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var question models.Question

	id,_ := strconv.Atoi(vars["id"])

	delete := question.Delete(id)
	fmt.Println(delete)
}