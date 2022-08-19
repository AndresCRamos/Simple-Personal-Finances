package routes

import (
	"encoding/json"
	"log"
	"net/http"

	bill "github.com/AndresCRamos/Simple-Personal-Finances/models/bills"
	earning "github.com/AndresCRamos/Simple-Personal-Finances/models/earning"
	incomesource "github.com/AndresCRamos/Simple-Personal-Finances/models/income_source"
	auth_user "github.com/AndresCRamos/Simple-Personal-Finances/pkg/auth/models/user"
	"github.com/AndresCRamos/Simple-Personal-Finances/pkg/utils"
	"github.com/gorilla/mux"
)

type path struct {
	Url    string `json:"path"`
	Method string `json:"method"`
}

func LoadRoutes() {
	var rootUrls = []utils.Route{
		{Url: "/", Handler: getPaths, Method: "GET"},
	}
	utils.AddRoutes(rootUrls)
	utils.AddRoutes(incomesource.SourcesRoutes)
	utils.AddRoutes(bill.BillsRoutes)
	utils.AddRoutes(earning.EarningsRoutes)
	utils.AddRoutes(auth_user.AuthUserRoutes)
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
