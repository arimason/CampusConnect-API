package utils

import "golang.org/x/crypto/bcrypt"

// gera um hash a partir de um campo string
func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// verifica se o campo é válido de acordo com o hash gerado
func ValidateHash(hash, origin string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(origin))
}
