package incomesource

import (
	"encoding/json"
	"net/http"

	auth_token "github.com/AndresCRamos/Simple-Personal-Finances/auth/models/token"
	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
	"github.com/emvi/null"
	"github.com/gorilla/mux"
)

func GetIncomeSourcesByUserID(w http.ResponseWriter, r *http.Request) {
	tokenObj, valid := auth_token.VerifyToken(w, r)
	if !valid {
		return
	}
	var incomeSources []IncomeSource
	var incomeSourcesGet []IncomeSourceGet
	if err := utils.Instance.Find(&incomeSources, "user_id = ?", tokenObj.User_id).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
	}
	for _, sourceItem := range incomeSources {
		incomeSourcesGet = append(incomeSourcesGet, IncomeSourceGet(sourceItem))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(incomeSourcesGet)
}

func SearchIncomeSourceByID(id string, user_id uint) (IncomeSource, bool, string) {
	var source IncomeSource
	err := utils.Instance.First(&source, "id = ? AND user_id = ?", id, user_id).Error
	found := true
	errorString := ""
	if err != nil {
		found = false
		errorString = err.Error()
	}
	return source, found, errorString
}

func GetIncomeSourcesByID(w http.ResponseWriter, r *http.Request) {
	tokenObj, valid := auth_token.VerifyToken(w, r)
	if !valid {
		return
	}
	sourceId := mux.Vars(r)["id"]
	source, found, err := SearchIncomeSourceByID(sourceId, tokenObj.User_id)
	if !found {
		utils.DisplaySearchError(w, r, "Sources", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(source)
	}
}

func GetIncomeSourcesDetailByID(w http.ResponseWriter, r *http.Request) {
	tokenObj, valid := auth_token.VerifyToken(w, r)
	if !valid {
		return
	}
	sourceId := mux.Vars(r)["id"]
	source, found, err := SearchIncomeSourceByID(sourceId, tokenObj.User_id)
	if !found {
		utils.DisplaySearchError(w, r, "Sources", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		incomeSourceDetail := IncomeSourceDetail(source)
		json.NewEncoder(w).Encode(&incomeSourceDetail)
	}
}

func CreateIncomeSource(w http.ResponseWriter, r *http.Request) {
	tokenObj, validToken := auth_token.VerifyToken(w, r)
	if !validToken {
		return
	}
	var incomeSource IncomeSource
	json.NewDecoder(r.Body).Decode(&incomeSource)
	incomeSource.User_id = null.NewInt64(int64(tokenObj.User_id), true)
	valid := utils.Validate(w, "Source", incomeSource)
	if !valid {
		return
	}

	if err := utils.Instance.Create(&incomeSource).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		incomeSourceGet := IncomeSourceGet(incomeSource)
		json.NewEncoder(w).Encode(&incomeSourceGet)
	}
}

func UpdateIncomeSource(w http.ResponseWriter, r *http.Request) {
	tokenObj, valid := auth_token.VerifyToken(w, r)
	if !valid {
		return
	}
	sourceId := mux.Vars(r)["id"]
	source, found, err := SearchIncomeSourceByID(sourceId, tokenObj.User_id)

	if !found {
		utils.DisplaySearchError(w, r, "Sources", err)
		return
	}
	json.NewDecoder(r.Body).Decode(&source)
	if !utils.Validate(w, "Source", source) {
		return
	} else if err := utils.Instance.Save(&source).Error; err != nil {
		utils.DisplaySearchError(w, r, "Sources", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(source)
	}
}

func DeleteIncomeSource(w http.ResponseWriter, r *http.Request) {
	tokenObj, valid := auth_token.VerifyToken(w, r)
	if !valid {
		return
	}
	sourceId := mux.Vars(r)["id"]
	source, found, err := SearchIncomeSourceByID(sourceId, tokenObj.User_id)
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
