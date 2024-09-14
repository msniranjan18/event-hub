package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWT_SECRET_KEY = "some_secret_key"

func GenerateToken(email string, userId int64) (string, error) {
	// setting the token expiration time to 2 hours
	expirationTime := time.Now().Add(time.Hour * 2)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    expirationTime.Unix(),
	})
	tokenString, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}
	if !parsedToken.Valid {
		return 0, errors.New("not valid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("not valid claims")
	}
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
