package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func ValidateHash(hash, origin string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(origin))
}
