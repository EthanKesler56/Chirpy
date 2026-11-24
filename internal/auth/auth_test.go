package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	UserID := uuid.New()
	secret := "Testing-Secret"
	jwtMakeTest, err := MakeJWT(UserID, secret, time.Hour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	jwtVal, err := ValidateJWT(jwtMakeTest, secret)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if jwtVal != UserID {
		t.Fatalf("got %v, want %v", jwtVal, UserID)
	}

}

func TestExpiredToken(t *testing.T) {
	UserID := uuid.New()
	Secret := "Testing-Secret"
	jwtExpiredTest, err := MakeJWT(UserID, Secret, -(time.Hour))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = ValidateJWT(jwtExpiredTest, Secret)
	if err == nil {
		t.Fatalf("expected an error for an expired token, got nil")
	}

}

func TestWrongSecret(t *testing.T) {
	userId := uuid.New()
	secret := "Testing-Secret"
	wrong_secret := "Wrong-Secret"
	JwtWrongSecretTest, err := MakeJWT(userId, secret, (time.Hour))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = ValidateJWT(JwtWrongSecretTest, wrong_secret)
	if err == nil {
		t.Fatalf("Wrong Secret!")
	}

}
