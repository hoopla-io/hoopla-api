package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hoopla/hoopla-api/app/config"
)

type UserClaims struct {
	UserID      uint   `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
}

func EncodeJWT(userId uint, phoneNumber string, expireAt int64) (string, error) {
	appCnf := config.NewAppConfig()

	claims := jwt.MapClaims{
		"iss":         "https://api.hoopla.uz",
		"aud":         "https://hoopla.uz",
		"iat":         time.Now().Unix(),
		"nbf":         time.Now().Unix(),
		"exp":         expireAt,
		"jti":         fmt.Sprintf("auth-hoopla-uz-%s", fmt.Sprintf("%v", time.Now().UnixNano())), // Unique ID
		"userId":      userId,
		"phoneNumber": phoneNumber,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := token.SignedString([]byte(appCnf.JwtKey))
	if err != nil {
		return "", err
	}

	return encodedToken, nil
}

func DecodeJWT(encodedToken string) (*UserClaims, error) {
	appCnf := config.NewAppConfig()
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(appCnf.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	userClaims := UserClaims{
		UserID:      uint(claims["userId"].(float64)),
		PhoneNumber: fmt.Sprintf("%v", claims["phoneNumber"]),
	}

	return &userClaims, nil
}
