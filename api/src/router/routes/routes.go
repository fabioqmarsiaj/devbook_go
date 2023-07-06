package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes represents all API routes
type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Put all routes inside the router
func Config(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}