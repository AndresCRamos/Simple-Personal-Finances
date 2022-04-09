package incomesource

import (
	"encoding/json"
	"net/http"

	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
	"github.com/gorilla/mux"
)

func GetIncomeSourcesByUserID(w http.ResponseWriter, r *http.Request) {
	var incomeSources []IncomeSource
	if err := utils.Instance.Find(&incomeSources).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(incomeSources)
}

func GetIncomeSourcesByID(w http.ResponseWriter, r *http.Request) {
	var source IncomeSource
	sourceId := mux.Vars(r)["id"]
	if err := utils.Instance.First(&source, sourceId).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(source)
	}
}

func CreateIncomeSource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var incomeSource IncomeSource
	json.NewDecoder(r.Body).Decode(&incomeSource)
	errorList, valid := utils.Validate(incomeSource)
	if !valid {
		utils.DisplayFieldErrors(w, r, "Source", errorList)
	} else if err := utils.Instance.Create(&incomeSource).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(incomeSource)
	}
}
