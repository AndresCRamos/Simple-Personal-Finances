package auth_user

import "github.com/AndresCRamos/Simple-Personal-Finances/utils"

var AuthUserRoutes = []utils.Route{
	{Url: "/register", Handler: Register, Method: "POST"},
}
