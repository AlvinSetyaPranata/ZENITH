package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessSecret  = []byte("access_secret_example")
	refreshSecret = []byte("refresh_secret_example")
)

func GenerateToken(userID string, role string) (string, string) {

	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessSignedToken, err := accessToken.SignedString(accessSecret)
	if err != nil {
		return "", ""
	}

	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshSignedToken, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		return "", ""
	}

	return accessSignedToken, refreshSignedToken
}

func ParseToken(tokenStr string, isAccessToken bool) (string, string, error) {
	secret := accessSecret
	if !isAccessToken {
		secret = refreshSecret
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil || !token.Valid {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", fmt.Errorf("invalid claims")
	}

	userID, _ := claims["user_id"].(string)
	role, _ := claims["role"].(string)

	return userID, role, nil
}
