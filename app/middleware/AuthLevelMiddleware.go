package middleware

import (
	"encoding/json"
	"net/http"

	// "strings"

	"../helpers"
	"../models"
)

type AuthLevelMiddleware struct {
}

func (q AuthLevelMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var route helpers.Route

		if route.IsPublicRoute(r.URL.Path) {
			next.ServeHTTP(w, r)

			return
		}

		sessionToken := r.Header.Get("SessionToken")

		var user models.User

		authUser, _ := user.GetUserByField("session_token", sessionToken)

		if !user.IsAuthorized(authUser, r.URL.Path, r.Method) {
			res, _ := json.Marshal("You are not authorized to visit this link.")
			w.Write(res)

			return
		}

		next.ServeHTTP(w, r)
	})
}
