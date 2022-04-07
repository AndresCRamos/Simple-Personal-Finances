package server

import (
	"fmt"
	"net/http"
)

func loadRoutes() {
	AddRoute("/api/v1", load, "GET")
}

func AddRoute(routeString string, handler func(http.ResponseWriter, *http.Request), method string) {
	router.HandleFunc(routeString, handler).Methods(method)
}

func load(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to this simple app")
}
