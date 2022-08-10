package server

import (
	"log"
	"net/http"

	"github.com/AndresCRamos/Simple-Personal-Finances/pkg/utils"
)

func Serve(port string) {

	log.Println("Serving...")
	log.Println("Now listening on localhost" + port)
	log.Fatal(http.ListenAndServe(port, utils.Router))
}
