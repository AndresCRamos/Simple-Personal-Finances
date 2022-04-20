package auth_user

import (
	"encoding/json"
	"net/http"
	"time"

	auth_token "github.com/AndresCRamos/Simple-Personal-Finances/auth/models/token"
	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
	"golang.org/x/crypto/bcrypt"
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
	} else if err := utils.Instance.Create(user.Parse()).Error; err != nil {
		utils.DisplaySearchError(w, r, "Users", err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user.Parse())
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	var userLog UserRegister
	if Jsonerr := json.NewDecoder(r.Body).Decode(&userLog); Jsonerr != nil {
		utils.DisplaySearchError(w, r, "Login", Jsonerr.Error())
		return
	}

	if err := utils.Instance.First(&user, "email = ?", userLog.Email).Error; err != nil {
		utils.DisplaySearchError(w, r, "Login", err.Error())
		return
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(userLog.Password.String)); err != nil {
		utils.DisplaySearchError(w, r, "Login", err.Error())
		return
	}
	bToken, tokenErr := auth_token.GenerateToken()
	if tokenErr != nil {
		utils.DisplaySearchError(w, r, "Login", tokenErr.Error())
		return
	}
	token := auth_token.Token{
		Token:     bToken,
		User_id:   user.ID,
		ExpiresAt: time.Now().Add(1 * time.Minute),
	}
	utils.Instance.Create(&token)
	json.NewEncoder(w).Encode(&struct {
		User  User
		Token auth_token.Token
	}{
		User:  user,
		Token: token,
	},
	)
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Token")
	valid := auth_token.DeleteToken(w, r, token)
	if !valid {
		return
	}
	json.NewEncoder(w).Encode(&struct {
		Message string
	}{
		"Logout successfull",
	})

}
