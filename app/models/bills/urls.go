package bill

import "github.com/AndresCRamos/Simple-Personal-Finances/pkg/utils"

var BillsRoutes = []utils.Route{
	{Url: "/api/v1/bill", Handler: GetBillsByUserID, Method: "GET"},
	{Url: "/api/v1/bill/{id}", Handler: GetBillByID, Method: "GET"},
	{Url: "/api/v1/bill", Handler: CreateBill, Method: "POST"},
	{Url: "/api/v1/bill/{id}", Handler: UpdateBill, Method: "PUT"},
	{Url: "/api/v1/bill/{id}", Handler: DeleteBill, Method: "DELETE"},
}
