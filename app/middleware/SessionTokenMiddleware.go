package middleware

import (
	"net/http"
	"encoding/json"
	"strings"

	"../models"
)


type TokenSessionMiddleware struct {
}

func (q TokenSessionMiddleware) Handle(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isRouteExcluded(r.URL.Path) {
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

func isRouteExcluded(path string) bool {
	uriSegments := strings.Split(path, "/")

	excludedRoutes := [4]string{"login", 
								"logout", 
								"register", 
								"password"}

	for _, n := range excludedRoutes {
		if uriSegments[1] == n {
			return true
		}
	}

	return false
}