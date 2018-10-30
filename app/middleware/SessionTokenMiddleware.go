package middleware

import (
	"encoding/json"
	"net/http"

	"../helpers"
	"../models"
)

type TokenSessionMiddleware struct {
}

func (q TokenSessionMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var route helpers.Route

		if route.IsPublicRoute(r.URL.Path) {
			next.ServeHTTP(w, r)

			return
		}

		sessionToken := r.Header.Get("SessionToken")

		var user models.User

		if !user.IsAuthenticated(sessionToken) {
			res, _ := json.Marshal("You are not logged in.")
			w.Write(res)

			return
		}

		next.ServeHTTP(w, r)
	})
}
