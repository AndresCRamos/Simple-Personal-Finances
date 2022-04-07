package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter().StrictSlash(true)

func Serve(port string) {
	loadRoutes()
	log.Println("Serving...")
	log.Println("Now listening on localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
