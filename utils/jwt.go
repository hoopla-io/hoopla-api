package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type JWTConfig struct {
	ISS string
	AUD string
	Key string
}

func EncodeJWT(userId uint, phoneNumber string, expireAt int64) (string, error) {
	godotenv.Load("." + "/.env")
	conf := viper.New()

	conf.AutomaticEnv()

	claims := jwt.MapClaims{
		"iat":         time.Now().Unix(),
		"nbf":         time.Now().Unix(),
		"exp":         expireAt,
		"jti":         fmt.Sprintf("auth-qahvzor-uz-%s", fmt.Sprintf("%v", time.Now().UnixNano())), // Unique ID
		"userId":      userId,
		"phoneNumber": phoneNumber,
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	encodedToken, err := token.SignedString([]byte(conf.GetString("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return encodedToken, nil
}
