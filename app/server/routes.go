package server

import (
	"encoding/json"
	"log"
	"net/http"

	incomesource "github.com/AndresCRamos/Simple-Personal-Finances/models/income_source"
	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
	"github.com/gorilla/mux"
)

type path struct {
	Url    string `json:"path"`
	Method string `json:"method"`
}

func loadRoutes() {
	var rootUrls = []utils.Route{
		{Url: "/", Handler: getPaths, Method: "GET"},
	}
	utils.AddRoutes(rootUrls)
	utils.AddRoutes(incomesource.SourcesRoutes)
}

func getPaths(w http.ResponseWriter, r *http.Request) {
	var paths []path
	err := utils.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathString, _ := route.GetPathTemplate()
		method, _ := route.GetMethods()
		newPath := path{Url: pathString, Method: method[0]}
		paths = append(paths, newPath)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(paths)
}
