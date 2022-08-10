package main

import (
	"github.com/AndresCRamos/Simple-Personal-Finances/cmd/server"
	"github.com/AndresCRamos/Simple-Personal-Finances/cmd/server/settings"
	"github.com/AndresCRamos/Simple-Personal-Finances/pkg/db"
)

func main() {
	config := settings.LoadConfig()
	db.Connect(config.GetDB())
	db.Migrate()
	server.Serve(config.GetPort())
}
