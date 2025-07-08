package middleware

import (
	"context"
	"crypto/rsa"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rickyhuang08/99-backend-exercise/helpers"
	"github.com/rickyhuang08/99-backend-exercise/pkg/auth"
)

type AuthModule struct{
	Time helpers.TimeProvider
}

func NewAuthModule(time helpers.TimeProvider) *AuthModule {
	return &AuthModule{
		Time: time,
	}
}

// AuthMiddleware verifies JWT tokens
func (m *AuthModule) AuthMiddleware(publicKey *rsa.PublicKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := m.ValidateJWT(tokenString, publicKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Store user info in context
		ctx := context.WithValue(c.Request.Context(), auth.UserKey, claims)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// ValidateJWT validates a JWT token
func (m *AuthModule) ValidateJWT(tokenString string, publicKey *rsa.PublicKey) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return publicKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}

	// Check expiration
	if exp, ok := claims["exp"].(float64); ok {
		if int64(exp) < m.Time.Now().Unix() {
			return nil, jwt.ErrTokenExpired
		}
	}

	return claims, nil
}