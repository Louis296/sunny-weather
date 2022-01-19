package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/louis296/sunny-weather/dao/model"
	"time"
)

var Secret string

type Claims struct {
	Email string
	Name  string
	jwt.StandardClaims
}

func GenerateToken(user model.User) (string, error) {
	claims := Claims{
		Email: user.Email,
		Name:  user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Subject:   "sunny_weather",
			Issuer:    "sunny_weather_louis296",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
