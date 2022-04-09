package incomesource

import (
	"encoding/json"
	"net/http"

	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
	"github.com/gorilla/mux"
)

func GetIncomeSourcesByUserID(w http.ResponseWriter, r *http.Request) {
	var incomeSources []IncomeSource
	var incomeSourcesGet []IncomeSourceGet
	if err := utils.Instance.Find(&incomeSources).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
	}
	for _, sourceItem := range incomeSources {
		incomeSourcesGet = append(incomeSourcesGet, IncomeSourceGet(sourceItem))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(incomeSourcesGet)
}

func SearchIncomeSourceByID(id string) (IncomeSource, bool, string) {
	var source IncomeSource
	err := utils.Instance.First(&source, id).Error
	found := true
	errorString := ""
	if err != nil {
		found = false
		errorString = err.Error()
	}
	return source, found, errorString
}

func GetIncomeSourcesByID(w http.ResponseWriter, r *http.Request) {
	sourceId := mux.Vars(r)["id"]
	source, found, err := SearchIncomeSourceByID(sourceId)
	if !found {
		utils.DisplaySearchError(w, r, "Sources", err)
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
		incomeSourceGet := IncomeSourceGet(incomeSource)
		json.NewEncoder(w).Encode(&incomeSourceGet)
	}
}

func UpdateIncomeSource(w http.ResponseWriter, r *http.Request) {
	var source IncomeSource
	sourceId := mux.Vars(r)["id"]
	if err := utils.Instance.First(&source, sourceId).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
		return
	}
	json.NewDecoder(r.Body).Decode(&source)
	if errorList, valid := utils.Validate(source); !valid {
		utils.DisplayFieldErrors(w, r, "Source", errorList)
	} else if err := utils.Instance.Save(&source).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(source)
	}
}

func DeleteIncomeSource(w http.ResponseWriter, r *http.Request) {
	sourceId := mux.Vars(r)["id"]
	source, found, err := SearchIncomeSourceByID(sourceId)
	if !found {
		utils.DisplaySearchError(w, r, "Sources", err)
	} else if err := utils.Instance.Delete(&source).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Deleted")
	}
}
