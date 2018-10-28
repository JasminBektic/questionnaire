package auth

import (
	// "fmt"
	"net/http"
	// "io/ioutil"
	// "encoding/json"
	// // "os"
	// "../../models"
)


type LoginController struct {
}

func (l LoginController) Login(w http.ResponseWriter, r *http.Request) {
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // TODO: Field validator - username and pass

	// var user models.User
	// var res []byte

	// err = json.Unmarshal(body, &user)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// m := map[string]string{"email":user.Email};

	// getUser, err := user.LoginFindByField(m)
	// if err == nil {
	// 	res, _ = json.Marshal("User already exists.")
	// 	w.Write(res)

	// 	return
	// }

	// fmt.Println(getUser)

	// res, _ = json.Marshal("test")

	// w.Write(res)
}

func (l LoginController) Logout(w http.ResponseWriter, r *http.Request) {
	
}
