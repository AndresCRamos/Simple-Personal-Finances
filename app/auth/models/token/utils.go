package auth_token

import (
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

func SearchToken(btoken string) (Token, bool) {
	var token Token
	valid := true
	if err := utils.Instance.First(&token, "token = ?", btoken).Error; err != nil {
		valid = false
	}
	return token, valid
}
