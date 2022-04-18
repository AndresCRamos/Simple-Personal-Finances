package auth_user

import (
	"encoding/json"
	"net/http"

	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user UserRegister
	errJson := json.NewDecoder(r.Body).Decode(&user)
	if errJson != nil {
		utils.DisplaySearchError(w, r, "Users", errJson.Error())
		return
	}
	valid := utils.Validate(w, "Login", user)
	if !valid {
		return
		// } else if err := utils.Instance.Create(user.Parse()).Error; err != nil {
		// 	utils.DisplaySearchError(w, r, "Users", err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user.Parse())
	}
}

func Login() {}
