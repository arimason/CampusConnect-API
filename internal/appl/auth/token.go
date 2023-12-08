package authappl

import (
	"campusconnect-api/configs"
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtToken struct {
	UserID     string
	Email      string
	Expiration time.Time
}

func createToken(id, email, secret string) (string, error) {
	// geração do token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"email":   email,
		// "exp":   time.Now().Add(300 * time.Second).Unix(),
	})
	// assina o token com a secret
	jwt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	// sucesso
	return jwt, nil
}

func verifierToken(ctx context.Context) (*jwt.Token, error) {
	// obtendo configs
	cfg, err := configs.LoadConfigs("./configs/app.yaml")
	if err != nil {
		return nil, errors.New("erro ao obter configs")
	}
	token := ctx.Value("JWT")
	// verifica a assinatura do token
	tk, err := jwt.Parse(token.(string), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("assinatura do token inválida!")
		}
		return []byte(cfg.JWTSecret), nil
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

func tokenValues(ctx context.Context) (*jwtToken, error) {
	tk, err := verifierToken(ctx)
	if err != nil {
		return nil, err
	}
	claims := tk.Claims.(jwt.MapClaims)
	token := &jwtToken{
		UserID: claims["user_id"].(string),
		Email:  claims["email"].(string),
	}
	return token, nil
}
