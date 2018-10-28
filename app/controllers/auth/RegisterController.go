package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../../models"
	"github.com/gorilla/mux"
)


type RegisterController struct {
}

/*
 *  First step in user registration process
 */
func (reg RegisterController) Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var user models.User
	var res []byte

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	m := map[string]string{"email":user.Email};

	getUser, err := user.FindByFields(m)
	if err == nil {
		res, _ = json.Marshal("User already exists.")
		w.Write(res)

		return
	}

	inserted := user.Insert(getUser)

	res, _ = json.Marshal("register/finish/" + inserted.Email + "/" + inserted.Token + "")

	w.Write(res)
}

/*
 *  Final step in user registration process
 */
func (reg RegisterController) FinishRegistration(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var user models.User
	var res []byte

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	vars := mux.Vars(r)

	user.Email = vars["email"]
	user.Token = vars["token"]

	m := map[string]string{"email":user.Email, "token":user.Token};
	
	user, err = user.FindByFields(m)
	if err != nil {
		res, _ = json.Marshal("Invalid url or you are already registered.")
		w.Write(res)

		return
	}
	fmt.Println(user)
	user.Update(user)

	res, _ = json.Marshal("Registration completed.")

	w.Write(res)
}
