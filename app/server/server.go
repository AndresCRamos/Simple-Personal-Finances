package server

import (
	"log"
	"net/http"

	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
)

func Serve(port string) {
	loadRoutes()
	log.Println("Serving...")
	log.Println("Now listening on localhost" + port)
	log.Fatal(http.ListenAndServe(port, utils.Router))
}
