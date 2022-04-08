package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter().StrictSlash(true)

type Route struct {
	Url     string
	Handler func(w http.ResponseWriter, r *http.Request)
	Method  string
}

type errorBuilder struct {
	Source  string `json:"source"`
	Message string `json:"message"`
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func AddRoutes(routeSlice []Route) error {
	for _, route := range routeSlice {
		Router.HandleFunc(route.Url, route.Handler).Methods(route.Method)
	}
	return nil
}

func DisplayError(w http.ResponseWriter, r *http.Request, source string, message string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorBuilder{Source: source, Message: message})

}
