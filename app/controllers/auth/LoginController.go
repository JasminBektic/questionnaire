package auth

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"../../models"

	"golang.org/x/crypto/bcrypt"
)


type LoginController struct {
}

/*
 *  Login process
 */
func (l LoginController) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: Field validator - username and pass

	var user models.User
	var res []byte

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	getUser, err := user.GetUserByField("username", user.Username)
	if err != nil {
		res, _ = json.Marshal("User with that username does not exists.")
		w.Write(res)

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(user.Password))
	if err != nil {
		res, _ = json.Marshal("Wrong password.")
		w.Write(res)

		return
	}

	sessionToken := user.SetAuthToken(getUser, true)

	res, _ = json.Marshal("Session token: " + sessionToken)
	w.Write(res)
}

/*
 *  Logout process
 */
func (l LoginController) Logout(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var res []byte

	authUser, err := user.GetUserByField("session_token", r.Header.Get("SessionToken"))
	if err != nil {
		res, _ = json.Marshal("You are not logged in.")
		w.Write(res)

		return
	}

	user.SetAuthToken(authUser, false)
	
	res, _ = json.Marshal("You logged out.")
	w.Write(res)
}
