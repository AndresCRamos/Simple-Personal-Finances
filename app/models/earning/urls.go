package earning

import "github.com/AndresCRamos/Simple-Personal-Finances/utils"

var EarningsRoutes = []utils.Route{
	{Url: "/api/v1/earning", Handler: GetEarningsByUserID, Method: "GET"},
	{Url: "/api/v1/earning/{id}", Handler: GetEarningByID, Method: "GET"},
	{Url: "/api/v1/earning", Handler: CreateEarning, Method: "POST"},
	{Url: "/api/v1/earning/{id}", Handler: UpdateEarning, Method: "PUT"},
	{Url: "/api/v1/earning/{id}", Handler: DeleteEarning, Method: "DELETE"},
}
