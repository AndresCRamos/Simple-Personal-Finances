package earning

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AndresCRamos/Simple-Personal-Finances/models/class"
	auth_token "github.com/AndresCRamos/Simple-Personal-Finances/pkg/auth/models/token"
	"github.com/AndresCRamos/Simple-Personal-Finances/pkg/utils"
	"github.com/emvi/null"
	"github.com/gorilla/mux"
)

func UpdateBalance(balance float64, source_id uint) {
	utils.Instance.Exec("UPDATE income_sources SET balance = balance + ? WHERE id = ?", balance, source_id)
}

func GetEarningsBySourceId(ID uint, user_id uint) []EarningList {
	var earningList []class.Earning
	var earningListDetail []EarningList
	utils.Instance.Find(&earningList, "income_source_id = ? AND user_id = ?", ID, user_id)
	for _, currentEarning := range earningList {
		earningListDetail = append(earningListDetail, EarningList(currentEarning))
	}
	return earningListDetail
}

func GetEarningsByUserID(w http.ResponseWriter, r *http.Request) {
	tokenObj, valid := auth_token.VerifyToken(w, r)
	if !valid {
		return
	}
	var Earnings []class.Earning
	var EarningsGet []EarningGet
	if err := utils.Instance.Find(&Earnings, "user_id = ?", tokenObj.User_id).Error; err != nil {
		utils.DisplaySearchError(w, r, "earnings", err.Error())
	}
	for _, earningItem := range Earnings {
		EarningsGet = append(EarningsGet, EarningGet(earningItem))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(EarningsGet)
}

func SearchEarningByID(id string, user_id uint) (class.Earning, bool, string) {
	var earning class.Earning
	err := utils.Instance.First(&earning, "user_id = ? AND id = ?", user_id, id).Error
	found := true
	errorString := ""
	if err != nil {
		found = false
		errorString = err.Error()
	}
	return earning, found, errorString
}

func GetEarningByID(w http.ResponseWriter, r *http.Request) {
	tokenObj, valid := auth_token.VerifyToken(w, r)
	if !valid {
		return
	}
	earningId := mux.Vars(r)["id"]
	earning, found, err := SearchEarningByID(earningId, tokenObj.User_id)
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
	tokenObj, tokenValid := auth_token.VerifyToken(w, r)
	if !tokenValid {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var earningData EarningCreate
	errJson := json.NewDecoder(r.Body).Decode(&earningData)
	if errJson != nil {
		utils.DisplaySearchError(w, r, "earnings", errJson.Error())
		return
	}
	earning := *earningData.Parse()
	earning.User_id = null.NewInt64(int64(tokenObj.User_id), true)
	valid := utils.Validate(w, "Source", earning)
	if !valid {
		return
	} else if err := utils.Instance.Create(&earning).Error; err != nil {
		utils.DisplaySearchError(w, r, "earnings", err.Error())
	} else {
		UpdateBalance(earning.Amount.Float64, uint(earning.Income_Source_id.Int64))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		EarningGet := EarningGet(earning)
		json.NewEncoder(w).Encode(&EarningGet)
	}
}

func UpdateEarning(w http.ResponseWriter, r *http.Request) {
	tokenObj, valid := auth_token.VerifyToken(w, r)
	if !valid {
		return
	}
	var earning class.Earning
	var earningData EarningCreate
	earningId := mux.Vars(r)["id"]
	earning, found, err := SearchEarningByID(earningId, tokenObj.User_id)
	if !found {
		utils.DisplaySearchError(w, r, "earnings", err)
		return
	}
	json.NewDecoder(r.Body).Decode(&earningData)
	earning = *earningData.Parse()
	if !utils.Validate(w, "Source", earning) {
		return
	} else if err := utils.Instance.Where("id = ?", earningId).Updates(&earning).Error; err != nil {
		utils.DisplaySearchError(w, r, "earnings", err.Error())
	} else {
		UpdateBalance(earning.Amount.Float64, uint(earning.Income_Source_id.Int64))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&earning)
	}
}

func DeleteEarning(w http.ResponseWriter, r *http.Request) {
	tokenObj, valid := auth_token.VerifyToken(w, r)
	if !valid {
		return
	}
	earningId := mux.Vars(r)["id"]
	earning, found, err := SearchEarningByID(earningId, tokenObj.ID)
	if !found {
		utils.DisplaySearchError(w, r, "earnings", err)
	} else if err := utils.Instance.Delete(&earning).Error; err != nil {
		utils.DisplaySearchError(w, r, "earnings", err.Error())
	} else {
		UpdateBalance(earning.Amount.Float64*-1, uint(earning.Income_Source_id.Int64))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Deleted")
	}
}
