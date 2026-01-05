package infra

import (
	"rest-fiber/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type GenerateTokenParams struct {
	ID    string
	Email string
}

type TokenService interface {
	GenerateToken(params *GenerateTokenParams, envStr string) (string, error)
}

type tokenServiceImpl struct {
}

func NewTokenService(env config.Env) TokenService {
	return &tokenServiceImpl{}
}

func (s *tokenServiceImpl) GenerateToken(params *GenerateTokenParams, envStr string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    params.ID,
		"email": params.Email,
		"exp":   expirationTime.Unix(),
		"iat":   time.Now().Unix(),
	})
	return token.SignedString([]byte(envStr))
}
