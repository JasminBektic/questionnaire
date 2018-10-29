package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../../models"
	"github.com/gorilla/mux"
)

type PasswordResetController struct {
}

/*
 *  First step in password reset process
 */
func (p PasswordResetController) ResetRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: Field validator - email

	var user models.User
	var res []byte

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	user.DeletePasswordReset(user.Email)

	inserted := user.InsertPasswordReset(user)

	res, _ = json.Marshal("password/reset/" + inserted.Email + "/" + inserted.Token + " is delivered to email. When you visit link, form is presented - here you will enter password")
	w.Write(res)
}

/*
 *  Final step in password reset process
 */
func (p PasswordResetController) ResetFinish(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: Field validator - password

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

	getPassResetReq, err := user.GetPasswordReset(user)
	if err != nil {
		res, _ = json.Marshal("An error occured. Send another password reset request.")
		w.Write(res)

		return
	}

	user.UpdatePassword(user)
	user.DeletePasswordReset(getPassResetReq.Email)

	res, _ = json.Marshal("Your password is changed.")
	w.Write(res)
}
