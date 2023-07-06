package router

import "github.com/gorilla/mux"

// Return a router with routes configured
func Generate() *mux.Router {
	return mux.NewRouter()
}