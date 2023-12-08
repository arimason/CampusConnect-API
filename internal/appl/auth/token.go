package authappl

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createToken(id, secret string) (string, error) {
	// geração do token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	// assina o token com a secret
	jwt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	// sucesso
	return jwt, nil
}
