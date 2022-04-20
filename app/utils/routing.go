package utils

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter().StrictSlash(true)

type Route struct {
	Url     string
	Handler func(w http.ResponseWriter, r *http.Request)
	Method  string
}

func AddRoutes(routeSlice []Route) error {
	for _, route := range routeSlice {
		Router.HandleFunc(route.Url, route.Handler).Methods(route.Method)
	}
	return nil
}
