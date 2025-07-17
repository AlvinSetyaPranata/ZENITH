package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessSecret  = []byte("access_secret_example")
	refreshSecret = []byte("refresh_secret_example")
)

func GenerateToken(userID string, role string) (string, string, int64) {

	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     ACCESS_TOKEN_EXPIRED_TIME,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessSignedToken, err := accessToken.SignedString(accessSecret)
	if err != nil {
		return "", "", 0
	}

	refresh_expired_time := REFRESH_TOKEN_EXPIRED_TIME

	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     refresh_expired_time,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshSignedToken, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		return "", "", 0
	}

	return accessSignedToken, refreshSignedToken, refresh_expired_time
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

func InvalidateToken(tokenStr string) string {
	return ""
}
