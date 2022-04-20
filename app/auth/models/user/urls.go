package auth_user

import "github.com/AndresCRamos/Simple-Personal-Finances/utils"

var AuthUserRoutes = []utils.Route{
	{Url: "/register", Handler: Register, Method: "POST"},
	{Url: "/login", Handler: Login, Method: "POST"},
	{Url: "/logout", Handler: LogOut, Method: "GET"},
}
