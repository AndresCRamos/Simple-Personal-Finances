package incomesource

import "github.com/AndresCRamos/Simple-Personal-Finances/utils"

var SourcesRoutes = []utils.Route{
	{Url: "/api/v1/source", Handler: GetIncomeSourcesByUserID, Method: "GET"},
	{Url: "/api/v1/source/{id}", Handler: GetIncomeSourcesByID, Method: "GET"},
}
