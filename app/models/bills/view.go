package bill

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
	"github.com/gorilla/mux"
)

func GetBillsBySourceId(ID uint) []BillList {
	var billList []Bill
	var billListDetail []BillList
	utils.Instance.Find(&billList, "income_source_id = ?", ID)
	for _, currentBill := range billList {
		billListDetail = append(billListDetail, BillList(currentBill))
	}
	return billListDetail
}

func GetBillsByUserID(w http.ResponseWriter, r *http.Request) {
	var Bills []Bill
	var BillsGet []BillGet
	if err := utils.Instance.Find(&Bills).Error; err != nil {
		utils.DisplaySearchError(w, r, "bills", err.Error())
	}
	for _, billItem := range Bills {
		BillsGet = append(BillsGet, BillGet(billItem))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BillsGet)
}

func SearchBillByID(id string) (Bill, bool, string) {
	var bill Bill
	err := utils.Instance.First(&bill, id).Error
	found := true
	errorString := ""
	if err != nil {
		found = false
		errorString = err.Error()
	}
	return bill, found, errorString
}

func GetBillByID(w http.ResponseWriter, r *http.Request) {
	billId := mux.Vars(r)["id"]
	bill, found, err := SearchBillByID(billId)
	if !found {
		utils.DisplaySearchError(w, r, "bills", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		res, _ := bill.MarshalJSON()
		fmt.Fprint(w, string(res))
	}
}

func CreateBill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var billData BillCreate
	errJson := json.NewDecoder(r.Body).Decode(&billData)
	if errJson != nil {
		utils.DisplaySearchError(w, r, "bills", errJson.Error())
		return
	}
	bill := *billData.Parse()
	valid := utils.Validate(w, "Source", bill)
	if !valid {
		return
	} else if err := utils.Instance.Create(&bill).Error; err != nil {
		utils.DisplaySearchError(w, r, "bills", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		BillGet := BillGet(bill)
		json.NewEncoder(w).Encode(&BillGet)
	}
}

func UpdateBill(w http.ResponseWriter, r *http.Request) {
	var bill Bill
	var billData BillCreate
	billId := mux.Vars(r)["id"]
	if err := utils.Instance.First(&bill, billId).Error; err != nil {
		utils.DisplaySearchError(w, r, "bills", err.Error())
		return
	}
	json.NewDecoder(r.Body).Decode(&billData)
	bill = *billData.Parse()
	if !utils.Validate(w, "Source", bill) {
		return
	} else if err := utils.Instance.Where("id = ?", billId).Updates(&bill).Error; err != nil {
		utils.DisplaySearchError(w, r, "bills", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&bill)
	}
}

func DeleteBill(w http.ResponseWriter, r *http.Request) {
	billId := mux.Vars(r)["id"]
	bill, found, err := SearchBillByID(billId)
	if !found {
		utils.DisplaySearchError(w, r, "bills", err)
	} else if err := utils.Instance.Delete(&bill).Error; err != nil {
		utils.DisplaySearchError(w, r, "bills", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Deleted")
	}
}
