package auth

import (
	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	claim, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}
	uuid_id, err := uuid.Parse(claim)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid_id, nil
}
