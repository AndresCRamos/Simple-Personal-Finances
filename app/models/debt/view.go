package debt

import (
	"encoding/json"
	"net/http"

	auth_token "github.com/AndresCRamos/Simple-Personal-Finances/auth/models/token"
	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
	"github.com/emvi/null"
)

func CreateDebt(w http.ResponseWriter, r *http.Request) {
	tokenObj, tokenValid := auth_token.VerifyToken(w, r)
	if !tokenValid {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var debtData DebtCreate
	errJson := json.NewDecoder(r.Body).Decode(&debtData)
	if errJson != nil {
		utils.DisplaySearchError(w, r, "debts", errJson.Error())
		return
	}
	debt := *debtData.Parse()
	debt.Lender_user_id = null.NewInt64(int64(tokenObj.User_id), true)
	valid := utils.Validate(w, "Source", debt)
	if !valid {
		return
	} else if err := utils.Instance.Create(&debt).Error; err != nil {
		utils.DisplaySearchError(w, r, "debts", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&debt)
	}
}
