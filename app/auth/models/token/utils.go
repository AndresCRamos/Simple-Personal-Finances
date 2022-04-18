package auth_token

import (
	"crypto/sha512"
	"math/rand"
)

func GenerateToken() ([]byte, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return b, err
	}
	hashFunc := sha512.New()
	hashFunc.Write(b)
	return hashFunc.Sum(nil), nil
}
