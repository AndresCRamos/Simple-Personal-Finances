package server

import (
	"fmt"
	"log"
	"net/http"
)

func Load(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to this simple app")
}

func Serve(port string) {
	http.HandleFunc("/", Load)
	log.Println("Serving...")
	log.Println("Now listening on localhost" + port)
	http.ListenAndServe(port, nil)
}
