package utils

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vladas9/backend-practice/internal/errors"
)

type Claims struct {
	UserID   string `json:"user_id"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID, userType string, jwtSecret []byte) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID:   userID,
		UserType: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.Internal(err)
	}

	return tokenString, nil
}

func ExtractUserIDFromToken(r *http.Request, jwtSecret []byte) (uuid.UUID, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return uuid.UUID{}, errors.Unauthorized("Authorisation header missing", nil)
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return uuid.UUID{}, errors.Unauthorized("Invalid Authorization header format. Expected 'Bearer <token>'", nil)
	}

	tokenString := parts[1]

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return uuid.UUID{}, errors.Unauthorized("Your session has expired", err)
	}

	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return uuid.UUID{}, errors.Unauthorized("Failed to parse user ID from token", err)
	}

	return userID, nil
}
