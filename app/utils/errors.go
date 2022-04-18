package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type simpleErrorAPIBuilder struct {
	Source string `json:"source"`
	Errors string `json:"errors"`
}

type listFieldErrorAPIBuilder struct {
	Source string       `json:"source"`
	Errors []FieldError `json:"errors"`
}

type FieldError struct {
	Field   string
	Message string
}

func DisplaySearchError(w http.ResponseWriter, r *http.Request, source string, errorList string) {
	log.Printf("404 %s, %v", source, errorList)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(simpleErrorAPIBuilder{Source: source, Errors: errorList})
}

func DisplayFieldErrors(w http.ResponseWriter, source string, errorList []FieldError) {
	log.Printf("404 %s, %v", source, errorList)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(listFieldErrorAPIBuilder{Source: source, Errors: errorList})
}
