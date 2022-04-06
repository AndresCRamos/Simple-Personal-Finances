package main

import (
	"fmt"
	"log"
	"net/http"
)

func Load(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func main() {
	http.HandleFunc("/", Load)
	log.Println("Serving...")
	http.ListenAndServe(":8080", nil)
}
