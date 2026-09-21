package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTempToken(secretTempKey []byte, userID string) (*string, *string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (data dalam token)
	expiredTime := time.Now().Add(time.Minute * 5).Unix()
	expiredTimeISO := time.Unix(expiredTime, 0).Format(time.RFC3339)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = userID
	claims["exp"] = expiredTime
	claims["type"] = "LOGIN_SELECTION"

	// Generate signed token string
	tokenString, err := token.SignedString(secretTempKey)
	if err != nil {
		return nil, nil, err
	}

	return &tokenString, &expiredTimeISO, nil
}

func GenerateToken(secretKey []byte, userID string, warehouseID string) (*string, *string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (data dalam token)
	expiredTime := time.Now().Add(time.Hour * 24).Unix()
	expiredTimeISO := time.Unix(expiredTime, 0).Format(time.RFC3339)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = userID
	claims["warehouse_id"] = warehouseID
	claims["exp"] = expiredTime
	claims["type"] = "ACCESS"

	// Generate signed token string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, nil, err
	}

	return &tokenString, &expiredTimeISO, nil
}
