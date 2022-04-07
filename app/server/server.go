package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Load(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to this simple app")
}

func LoadRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1", Load).Methods("GET")
}

func Serve(port string) {
	router := mux.NewRouter().StrictSlash(true)
	LoadRoutes(router)
	log.Println("Serving...")
	log.Println("Now listening on localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
