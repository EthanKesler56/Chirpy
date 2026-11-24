package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

type bearerTest struct {
	name       string
	authHeader string
	wantToken  string
	wantError  bool
}

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

func TestGetBearerToken(t *testing.T) {
	test := []bearerTest{{name: "valid", authHeader: "bearerabc123", wantToken: "", wantError: false}, {name: "invalid", authHeader: " ", wantToken: "SOmething", wantError: false}, {name: "valid", authHeader: "beardef456", wantToken: "poggers", wantError: false}}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			headers := http.Header{}
			if tc.authHeader != "" {
				headers.Set("Authorization", tc.authHeader)
			}
			BearerToken, err := GetBearerToken(headers)
			if (err != nil) != tc.wantError {
				t.Fatalf("got err %v, wantErr %v,", err, tc.wantError)
			}

		})
	}
}
