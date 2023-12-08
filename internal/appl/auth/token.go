package authappl

import (
	"context"
	"errors"
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

func verifierToken(ctx context.Context, secret string) (*jwt.Token, error) {
	token := ctx.Value("JWT")
	// verifica a assinatura do token
	tk, err := jwt.Parse(token.(string), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("assinatura do token inválida!")
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := tk.Claims.(jwt.MapClaims)
	if ok && tk.Valid {
		// obtenho a expiração com precisão de segundos
		exp := time.Unix(int64(claims["exp"].(float64)), 0)
		if time.Now().After(exp) {
			return nil, errors.New("token expirado!")
		}
	}
	return tk, nil
}
