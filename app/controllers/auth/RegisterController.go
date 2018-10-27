package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../../models"
	"github.com/gorilla/mux"
)

// "os"
// "../../models"

type RegisterController struct {
}

func (reg RegisterController) Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	insert := user.Insert(user)
	fmt.Println(insert)
}

func (reg RegisterController) FinishRegistration(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	vars := mux.Vars(r)

	user.Email = vars["email"]
	user.Token = vars["token"]

	update := user.Update(user)
	fmt.Println(update)
}
