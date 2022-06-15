package debt

import "github.com/AndresCRamos/Simple-Personal-Finances/utils"

var DebtsRoutes = []utils.Route{
	{Url: "/api/v1/debt", Handler: CreateDebt, Method: "POST"},
}
