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

var secretKey = []byte("capybara")       // TODO: add secret key via .env or some rotation
var claimSecretKey = []byte("capybara2") // TODO: add secret key via .env or some rotation

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

func CreateClaimToken(userData JWTUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"UserId":   userData.UserId,
			"Username": userData.Username,
			"Exp":      time.Now().Add(time.Minute * 5).Unix(),
		})
	tokenString, err := token.SignedString(claimSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateRefreshToken(userData JWTUser, isTimeBased bool, db *gorm.DB) (string, error) {
	var exp int64 = 0
	if isTimeBased {
		exp = time.Now().Add(time.Hour * 15).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"UserId":   userData.UserId,
			"Username": userData.Username,
			"Exp":      exp,
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	data := NewRefreshTokenDataDB{
		UserId:       userData.UserId,
		RefreshToken: tokenString,
	}
	err = PushRefreshTokenToDB(data, db)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string, isClaimRequest bool) (bool, error) {
	secret := secretKey
	if isClaimRequest {
		secret = claimSecretKey
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
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
var RefreshTokenIsNotValidDueToExpirationDate = errors.New("refresh token not valid due to expiration date")

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
	if payload.Exp > 0 {
		if payload.Exp < time.Now().Unix() {
			return payload, RefreshTokenIsNotValidDueToExpirationDate
		}
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
		Where("refresh_token = ?", token).
		Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

type NewRefreshTokenDataDB struct {
	UserId       string `json:"userId"`
	RefreshToken string `json:"refresh_token"`
}

func PushRefreshTokenToDB(data NewRefreshTokenDataDB, db *gorm.DB) error {

	result := db.Table("refreshTokens").Create(&data)
	return result.Error
}
