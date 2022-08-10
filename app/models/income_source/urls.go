package incomesource

import "github.com/AndresCRamos/Simple-Personal-Finances/pkg/utils"

var SourcesRoutes = []utils.Route{
	{Url: "/api/v1/source", Handler: GetIncomeSourcesByUserID, Method: "GET"},
	{Url: "/api/v1/source/{id}", Handler: GetIncomeSourcesByID, Method: "GET"},
	{Url: "/api/v1/source/{id}/bill", Handler: GetIncomeSourcesDetailByID, Method: "GET"},
	{Url: "/api/v1/source", Handler: CreateIncomeSource, Method: "POST"},
	{Url: "/api/v1/source/{id}", Handler: UpdateIncomeSource, Method: "PUT"},
	{Url: "/api/v1/source/{id}", Handler: DeleteIncomeSource, Method: "DELETE"},
}
