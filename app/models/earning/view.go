package earning

import (
	"encoding/json"
	"net/http"

	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
	"github.com/gorilla/mux"
)

func GetEarningsBySourceId(ID uint) []EarningList {
	var earningList []Earning
	var earningListDetail []EarningList
	utils.Instance.Find(&earningList, "income_source_id = ?", ID)
	for _, currentEarning := range earningList {
		earningListDetail = append(earningListDetail, EarningList(currentEarning))
	}
	return earningListDetail
}

func GetEarningsByUserID(w http.ResponseWriter, r *http.Request) {
	var Earnings []Earning
	var EarningsGet []EarningGet
	if err := utils.Instance.Find(&Earnings).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
	}
	for _, sourceItem := range Earnings {
		EarningsGet = append(EarningsGet, EarningGet(sourceItem))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(EarningsGet)
}

func SearchEarningByID(id string) (Earning, bool, string) {
	var source Earning
	err := utils.Instance.First(&source, id).Error
	found := true
	errorString := ""
	if err != nil {
		found = false
		errorString = err.Error()
	}
	return source, found, errorString
}

func GetEarningByID(w http.ResponseWriter, r *http.Request) {
	sourceId := mux.Vars(r)["id"]
	source, found, err := SearchEarningByID(sourceId)
	if !found {
		utils.DisplaySearchError(w, r, "Sources", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(source)
	}
}

func CreateEarning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Earning Earning
	json.NewDecoder(r.Body).Decode(&Earning)
	errorList, valid := utils.Validate(Earning)
	if !valid {
		utils.DisplayFieldErrors(w, r, "Source", errorList)
	} else if err := utils.Instance.Create(&Earning).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		EarningGet := EarningGet(Earning)
		json.NewEncoder(w).Encode(&EarningGet)
	}
}

func UpdateEarning(w http.ResponseWriter, r *http.Request) {
	var source Earning
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

func DeleteEarning(w http.ResponseWriter, r *http.Request) {
	sourceId := mux.Vars(r)["id"]
	source, found, err := SearchEarningByID(sourceId)
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
