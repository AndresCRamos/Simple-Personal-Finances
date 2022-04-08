package incomesource

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
	"github.com/gorilla/mux"
)

func checkIfExists(id string) (IncomeSource, bool) {
	var source IncomeSource
	utils.Instance.First(&source, id)
	fmt.Println(source)
	exists := false
	if source.ID != 0 {
		exists = true
	}
	return source, exists
}

func GetIncomeSourcesByUserID(w http.ResponseWriter, r *http.Request) {
	var incomeSources []IncomeSource
	utils.Instance.Find(&incomeSources)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(incomeSources)
}

func GetIncomeSourcesByID(w http.ResponseWriter, r *http.Request) {
	sourceId := mux.Vars(r)["id"]
	source, exists := checkIfExists(sourceId)
	if exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(source)
	} else {
		utils.DisplayError(w, r, "Sources", fmt.Sprintf("Source with id: %s does not exist", sourceId))
	}
}
