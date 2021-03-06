package models

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"strings"

	"../../db"
	"../helpers"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Type         int    `json:"type"`
	Token        string `json:"token"`
	SessionToken string `json:"session_token"`
}

/*
 *  Find user with multiple field
 */
func (u User) FindByFields(m map[string]string) (User, error) {
	db, err := db.Open()

	query := `SELECT 
				fullname, 
				email 
			FROM 
				users 
			WHERE `

	for k, v := range m {
		query += k + "='" + v + "' AND "
	}
	query = strings.TrimSuffix(query, " AND ")

	row := db.QueryRow(query)
	err = row.Scan(&u.Fullname, &u.Email)

	return u, err
}

/*
 *  Find user with specific field
 */
func (u User) GetUserByField(field string, value string) (User, error) {
	db, err := db.Open()

	row := db.QueryRow(`SELECT 
							id, 
							username, 
							fullname, 
							email, 
							type, 
							password 
						FROM 
							users 
						WHERE `+field+` = ?`, value)

	err = row.Scan(&u.Id, &u.Username, &u.Fullname, &u.Email, &u.Type, &u.Password)

	return u, err
}

/*
 *  User create
 */
func (u User) Insert(user User) User {
	db, _ := db.Open()

	user.Token = u.GenerateToken()

	db.QueryRow(`INSERT 
					INTO users (fullname, email, token) 
					VALUES (?, ?, ?)`, user.Fullname, user.Email, user.Token)

	return user
}

/*
 *  User update
 */
func (u User) Update(user User) sql.Result {
	db, _ := db.Open()

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(password)

	db.QueryRow(`UPDATE 
					users 
				SET username = ?, 
					password = ?, 
					token = NULL 
				WHERE email = ? AND token = ?`, user.Username, user.Password, user.Email, user.Token)

	return nil
}

/*
 *  Set new password
 */
func (u User) UpdatePassword(user User) sql.Result {
	db, _ := db.Open()

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(password)

	db.QueryRow(`UPDATE 
					users 
				SET password = ? 
				WHERE email = ?`, user.Password, user.Email)

	return nil
}

/*
 *  Get row from password_resets table
 */
func (u User) GetPasswordReset(user User) (User, error) {
	db, err := db.Open()

	row := db.QueryRow(`SELECT DISTINCT
							password_resets.email, 
							password_resets.token 
						FROM 
							password_resets
						INNER JOIN 
							users ON password_resets.email = users.email
						WHERE
							password_resets.email = ? AND password_resets.token = ?`, user.Email, user.Token)

	err = row.Scan(&u.Email, &u.Token)

	return u, err
}

/*
 *  Insert into password_resets table
 */
func (u User) InsertPasswordReset(user User) User {
	db, _ := db.Open()

	user.Token = u.GenerateToken()

	db.QueryRow(`INSERT 
					INTO password_resets (email, token) 
					VALUES (?, ?)`, user.Email, user.Token)

	return user
}

/*
 *  Delete resource from password_resets table
 */
func (u User) DeletePasswordReset(email string) sql.Result {
	db, _ := db.Open()

	db.QueryRow(`DELETE FROM password_resets WHERE email = ?`, email)

	return nil
}

/*
 *  Check if user session token is set
 */
func (u User) SetAuthToken(user User, activate bool) string {
	db, err := db.Open()
	if err != nil {
		panic(err)
	}

	query := `UPDATE users SET session_token = NULL WHERE id = ?`
	token := ""

	if activate {
		token = u.GenerateToken()
		query = `UPDATE users SET session_token = '` + token + `' WHERE id = ?`
	}

	row := db.QueryRow(query, user.Id)
	err = row.Scan(&u.SessionToken)

	return token
}

/*
 *  Check if user session token is set
 */
func (u User) IsAuthenticated(token string) bool {
	db, err := db.Open()

	row := db.QueryRow(`SELECT id FROM users WHERE session_token = ?`, token)
	err = row.Scan(&u.Id)
	if err != nil {
		return false
	}

	return true
}

/*
 *  Check if user has right route permissions
 */
func (u User) IsAuthorized(user User, path string, method string) bool {
	var r helpers.Route
	var routes map[string][]string

	switch user.Type {
	case 0:
		routes = r.AdminRoutes()
	case 1:
		routes = r.UserRoutes()
	}

	uriSegments := strings.Split(path, "/")

	for route, methods := range routes {
		if uriSegments[1] == route {
			for _, v := range methods {
				if method == v {
					return true
				}
			}
		}
	}

	return false
}

/*
 *  Generate random string
 */
func (u User) GenerateToken() string {
	b := make([]byte, 32)

	rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}
