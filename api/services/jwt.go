package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(payload string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = payload
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // expira em 24h
	tokenString, err := token.SignedString([]byte("remover-chave-provisoria-e-adicionar-uma-secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
