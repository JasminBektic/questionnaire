package helpers

import (
	// "net/http"
	// "encoding/json"
	"strings"

	// "../models"
)

type Route struct {
}

/*
 *  Check public route
 */
func (r Route) IsPublicRoute(path string) bool {
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

/*
 *  Admin route map
 */
func (r Route) AdminRoutes() map[string][]string {
	routes := map[string][]string{"survey":  []string{"GET", "POST", "PUT", "DELETE"}, 
								  "question":[]string{"GET", "POST", "PUT", "DELETE"},
								  "answer":  []string{"GET", "DELETE"}};

	return routes
}

/*
 *  User route map
 */
func (r Route) UserRoutes() map[string][]string {
	routes := map[string][]string{"survey":  []string{"GET"},
								  "answer":  []string{"POST"}};

	return routes
}