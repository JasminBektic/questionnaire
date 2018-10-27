package auth

import (
	// "fmt"
	"html/template"
	"net/http"
	// "os"
	// "../../models"
)

type LoginController struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (l LoginController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/auth/login.html")

		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, nil)
	} else {
		// var u models.User
		// u.Get()
		// fmt.Fprint(u.Get, "Homepage");
		// r.Form["username"]
		// r.Form["password"]
	}
	// fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
