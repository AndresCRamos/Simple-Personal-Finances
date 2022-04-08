package main

import (
	"github.com/AndresCRamos/Simple-Personal-Finances/db"
	"github.com/AndresCRamos/Simple-Personal-Finances/server"
	"github.com/AndresCRamos/Simple-Personal-Finances/settings"
)

func main() {
	config := settings.LoadConfig()
	db.Connect(config.GetDB())
	db.Migrate()
	server.Serve(config.GetPort())
}
