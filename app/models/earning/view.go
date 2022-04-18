package earning

import (
	"encoding/json"
	"fmt"
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
		utils.DisplaySearchError(w, r, "earnings", err.Error())
	}
	for _, earningItem := range Earnings {
		EarningsGet = append(EarningsGet, EarningGet(earningItem))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(EarningsGet)
}

func SearchEarningByID(id string) (Earning, bool, string) {
	var earning Earning
	err := utils.Instance.First(&earning, id).Error
	found := true
	errorString := ""
	if err != nil {
		found = false
		errorString = err.Error()
	}
	return earning, found, errorString
}

func GetEarningByID(w http.ResponseWriter, r *http.Request) {
	earningId := mux.Vars(r)["id"]
	earning, found, err := SearchEarningByID(earningId)
	if !found {
		utils.DisplaySearchError(w, r, "earnings", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		res, _ := earning.MarshalJSON()
		fmt.Fprint(w, string(res))
	}
}

func CreateEarning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var earningData EarningCreate
	errJson := json.NewDecoder(r.Body).Decode(&earningData)
	if errJson != nil {
		utils.DisplaySearchError(w, r, "earnings", errJson.Error())
		return
	}
	earning := *earningData.Parse()
	valid := utils.Validate(w, "Source", earning)
	if !valid {
		return
	} else if err := utils.Instance.Create(&earning).Error; err != nil {
		utils.DisplaySearchError(w, r, "earnings", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		EarningGet := EarningGet(earning)
		json.NewEncoder(w).Encode(&EarningGet)
	}
}

func UpdateEarning(w http.ResponseWriter, r *http.Request) {
	var earning Earning
	var earningData EarningCreate
	earningId := mux.Vars(r)["id"]
	if err := utils.Instance.First(&earning, earningId).Error; err != nil {
		utils.DisplaySearchError(w, r, "earnings", err.Error())
		return
	}
	json.NewDecoder(r.Body).Decode(&earningData)
	earning = *earningData.Parse()
	if !utils.Validate(w, "Source", earning) {
		return
	} else if err := utils.Instance.Where("id = ?", earningId).Updates(&earning).Error; err != nil {
		utils.DisplaySearchError(w, r, "earnings", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&earning)
	}
}

func DeleteEarning(w http.ResponseWriter, r *http.Request) {
	earningId := mux.Vars(r)["id"]
	earning, found, err := SearchEarningByID(earningId)
	if !found {
		utils.DisplaySearchError(w, r, "earnings", err)
	} else if err := utils.Instance.Delete(&earning).Error; err != nil {
		utils.DisplaySearchError(w, r, "earnings", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Deleted")
	}
}
