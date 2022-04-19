package bill

import (
	"encoding/json"
	"fmt"
	"net/http"

	auth_token "github.com/AndresCRamos/Simple-Personal-Finances/auth/models/token"
	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
	"github.com/emvi/null"
	"github.com/gorilla/mux"
)

func GetBillsBySourceId(ID uint, user_id uint) []BillList {
	var billList []Bill
	var billListDetail []BillList
	utils.Instance.Find(&billList, "income_source_id = ? AND user_id = ?", ID, user_id)
	for _, currentBill := range billList {
		billListDetail = append(billListDetail, BillList(currentBill))
	}
	return billListDetail
}

func GetBillsByUserID(w http.ResponseWriter, r *http.Request) {
	tokenObj, valid := auth_token.VerifyToken(w, r)
	if !valid {
		return
	}
	var Bills []Bill
	var BillsGet []BillGet
	if err := utils.Instance.Find(&Bills, "user_id = ?", tokenObj.User_id).Error; err != nil {
		utils.DisplaySearchError(w, r, "bills", err.Error())
	}
	for _, billItem := range Bills {
		BillsGet = append(BillsGet, BillGet(billItem))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BillsGet)
}

func SearchBillByID(id string, user_id uint) (Bill, bool, string) {
	var bill Bill
	err := utils.Instance.First(&bill, "user_id = ? AND id = ?", user_id, id).Error
	found := true
	errorString := ""
	if err != nil {
		found = false
		errorString = err.Error()
	}
	return bill, found, errorString
}

func GetBillByID(w http.ResponseWriter, r *http.Request) {
	tokenObj, valid := auth_token.VerifyToken(w, r)
	if !valid {
		return
	}
	billId := mux.Vars(r)["id"]
	bill, found, err := SearchBillByID(billId, tokenObj.User_id)
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
	tokenObj, tokenValid := auth_token.VerifyToken(w, r)

	if !tokenValid {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var billData BillCreate
	errJson := json.NewDecoder(r.Body).Decode(&billData)
	if errJson != nil {
		utils.DisplaySearchError(w, r, "bills", errJson.Error())
		return
	}
	bill := *billData.Parse()
	bill.User_id = null.NewInt64(int64(tokenObj.User_id), true)
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
	tokenObj, valid := auth_token.VerifyToken(w, r)
	if !valid {
		return
	}
	billId := mux.Vars(r)["id"]
	bill, found, err := SearchBillByID(billId, tokenObj.ID)
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
