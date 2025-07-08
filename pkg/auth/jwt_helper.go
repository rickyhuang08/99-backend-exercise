package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rickyhuang08/99-backend-exercise/helpers"
)

type JWTHelper struct {
	Time       helpers.TimeProvider
	PrivateKey string
}

func NewJWTHelper(time helpers.TimeProvider, privateKey string) *JWTHelper {
	return &JWTHelper{
		Time:       time,
		PrivateKey: privateKey,
	}
}

// GenerateJWT creates a signed JWT
func (m *JWTHelper) GenerateJWT(userID int) (string, error) {

	// Initialize usecases
	rsaPrivateKey, err := LoadPrivateKey(m.PrivateKey)

	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": m.Time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(rsaPrivateKey)
}
