package bill

import (
	"encoding/json"
	"net/http"

	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
	"github.com/gorilla/mux"
)

func GetBillsByUserID(w http.ResponseWriter, r *http.Request) {
	var Bills []Bill
	var BillsGet []BillGet
	if err := utils.Instance.Find(&Bills).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
	}
	for _, sourceItem := range Bills {
		BillsGet = append(BillsGet, BillGet(sourceItem))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BillsGet)
}

func SearchBillByID(id string) (Bill, bool, string) {
	var source Bill
	err := utils.Instance.First(&source, id).Error
	found := true
	errorString := ""
	if err != nil {
		found = false
		errorString = err.Error()
	}
	return source, found, errorString
}

func GetBillByID(w http.ResponseWriter, r *http.Request) {
	sourceId := mux.Vars(r)["id"]
	source, found, err := SearchBillByID(sourceId)
	if !found {
		utils.DisplaySearchError(w, r, "Sources", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(source)
	}
}

func CreateBill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Bill Bill
	json.NewDecoder(r.Body).Decode(&Bill)
	errorList, valid := utils.Validate(Bill)
	if !valid {
		utils.DisplayFieldErrors(w, r, "Source", errorList)
	} else if err := utils.Instance.Create(&Bill).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		BillGet := BillGet(Bill)
		json.NewEncoder(w).Encode(&BillGet)
	}
}

func UpdateBill(w http.ResponseWriter, r *http.Request) {
	var source Bill
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

func DeleteBill(w http.ResponseWriter, r *http.Request) {
	sourceId := mux.Vars(r)["id"]
	source, found, err := SearchBillByID(sourceId)
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
