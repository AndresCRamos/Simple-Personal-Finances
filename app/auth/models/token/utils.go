package auth_token

import (
	"net/http"

	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken() (string, error) {
	b := make([]byte, 32)
	tokenStr, err := bcrypt.GenerateFromPassword(b, 12)
	if err != nil {
		return "", err
	}
	return string(tokenStr), nil
}

func VerifyToken(w http.ResponseWriter, r *http.Request) (Token, bool) {
	token := r.Header.Get("Token")
	if token == "" {
		http.Error(w, "Undefined token", http.StatusUnauthorized)
		return Token{}, false
	}
	tokenObj, valid := searchToken(token)
	if !valid {
		http.Error(w, "Cant find this token", http.StatusUnauthorized)
		return Token{}, false
	}
	return tokenObj, true
}

func DeleteToken(w http.ResponseWriter, r *http.Request, token string) bool {
	tokenObj, valid := searchToken(token)
	if !valid {
		utils.DisplaySearchError(w, r, "Log out", "Cant find this token")
		return false
	}
	if !validateToken(tokenObj) {
		utils.DisplaySearchError(w, r, "Log out", "This token has expired")
		return false
	}
	if err := utils.Instance.Delete(&tokenObj).Error; err != nil {
		utils.DisplaySearchError(w, r, "Logout", err.Error())
		return false
	}
	return true
}

func searchToken(btoken string) (Token, bool) {
	var token Token
	valid := true
	if err := utils.Instance.First(&token, "token = ?", btoken).Error; err != nil {
		valid = false
	}
	return token, valid
}

func validateToken(token Token) bool {
	// return time.Now().Before(token.ExpiresAt)
	return true
}
