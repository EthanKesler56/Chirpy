package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {

	jwt_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{Issuer: "chirpy-access", IssuedAt: jwt.NewNumericDate(time.Now().UTC()), ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn).UTC()), Subject: userID.String()})
	token, err := jwt_token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return token, nil

}
