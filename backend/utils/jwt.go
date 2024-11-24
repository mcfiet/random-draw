package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var secretkey = []byte("Xj9v!@BzS#l2Fg$7HtUv5R*Lp8MfYqZ0WaKdQr1NxOiV$JpL")

func GenerateToken(userID uuid.UUID, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretkey)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing Method")
		}
		return secretkey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}

func GetUserIDFromToken(tokenString string) (uuid.UUID, error) {
	claims, err := VerifyToken(tokenString)
	if err != nil {
		return uuid.UUID{}, err
	}
	userID, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return uuid.UUID{}, err
	}
	return userID, nil
}

func ExtractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Authorization header is missing")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("Authorization header is not a bearer token")
	}
	return strings.TrimPrefix(authHeader, "Bearer "), nil
}
