package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("capybara") // TODO: add secret key via .env or some rotation

func CreateToken(userData JWTUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"UserId":   userData.UserId,
			"Username": userData.Username,
			"Exp":      time.Now().Add(time.Minute * 15).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {

		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return false, nil
		}
		return false, err
	}

	if !token.Valid {
		return false, nil
	}
	tokenData, err := DecodeBearer(tokenString)
	if err != nil {
		return false, err
	}

	if tokenData.Exp < time.Now().Unix() {
		return false, nil
	}

	return true, nil
}

func DecodeBearer(tokenString string) (JWTPayload, error) {
	splitToken := strings.Split(tokenString, ".")
	if len(splitToken) != 3 {
		return JWTPayload{}, fmt.Errorf("invalid token format")
	}

	payloadSegment := splitToken[1]
	payloadBytes, err := base64.RawURLEncoding.DecodeString(payloadSegment)
	if err != nil {
		return JWTPayload{}, fmt.Errorf("failed to decode payload: %v", err)
	}

	var payload JWTPayload
	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		return JWTPayload{}, fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	return payload, nil
}

var RefreshTokenNotInDbError = errors.New("refresh token not found in database")

func VerifyRefreshToken(tokenString string, db *gorm.DB) (JWTPayload, error) {

	splitToken := strings.Split(tokenString, ".")
	if len(splitToken) != 3 {
		return JWTPayload{}, fmt.Errorf("invalid token format")
	}

	payloadSegment := splitToken[1]
	payloadBytes, err := base64.RawURLEncoding.DecodeString(payloadSegment)
	if err != nil {
		return JWTPayload{}, fmt.Errorf("failed to decode payload: %v", err)
	}

	var payload JWTPayload
	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		return JWTPayload{}, fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	if payload.Exp < time.Now().Unix() {
		return payload, nil
	}

	inDB, err := VerifyRefreshTokenInDB(tokenString, payload.UserId, db)
	if err != nil {
		return payload, err
	}
	if !inDB {
		return payload, RefreshTokenNotInDbError
	}

	return payload, nil
}

func VerifyRefreshTokenInDB(token string, userId string, db *gorm.DB) (bool, error) {
	var count int64
	err := db.Table("refreshTokens").
		Where("user_id = ?", userId).
		Where("refreshToken = ?", token).
		Where("refreshToken = ?", token).
		Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}
