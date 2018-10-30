package main

import (
	"net/http"

	"./app/controllers"
	"./app/controllers/auth"
	"./app/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	var l auth.LoginController
	var r auth.RegisterController
	var pr auth.PasswordResetController
	var q controllers.QuestionController
	var s controllers.SurveyController
	var a controllers.AnswerController
	var ts middleware.TokenSessionMiddleware
	var al middleware.AuthLevelMiddleware

	// Middleware
	router.Use(ts.Handle)
	router.Use(al.Handle)

	// Routes
	router.HandleFunc("/login", l.Login).Methods("POST")
	router.HandleFunc("/logout", l.Logout).Methods("GET")
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

	router.HandleFunc("/answer", a.Insert).Methods("POST")

	http.ListenAndServe(":8000", router)
}
