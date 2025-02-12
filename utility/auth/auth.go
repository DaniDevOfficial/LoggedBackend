package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"loggedin/utility/jwt"
)

func GetJWTTokenFromHeader(c *gin.Context) (string, error) {
	jwtString := c.Request.Header.Get("Authorization")
	if jwtString == "" {
		return "", fmt.Errorf("missing authorization header")
	}
	return jwtString, nil
}

func GetRefreshTokenFromHeader(c *gin.Context) (string, error) {
	refreshToken := c.Request.Header.Get("RefreshToken")
	if refreshToken == "" {
		return "", fmt.Errorf("missing refresh token header")
	}
	return refreshToken, nil
}

// GetJWTPayloadFromHeader extracts the JWT payload from the Authorization header of an HTTP request.
// It first retrieves the JWT token from the header, verifies the token, and then decodes the payload.
//
// Parameters:
//
//	r (*http.Request): The HTTP request containing the Authorization header with the JWT token.
//
// Returns:
//
//	(jwt.JWTPayload, error): Returns the decoded JWT payload if successful, otherwise returns an error.
func GetJWTPayloadFromHeader(c *gin.Context, db *gorm.DB) (jwt.JWTPayload, error) {
	jwtToken, err := GetJWTTokenFromHeader(c)
	var jwtData jwt.JWTPayload
	if err != nil {
		return jwtData, err
	}
	valid, err := jwt.VerifyToken(jwtToken)
	if err != nil {
		return jwtData, err
	}
	if !valid {
		jwtData, newJwtToken, err := CreateNewTokenWithRefreshToken(c, db)
		log.Println(jwtData, newJwtToken, err)
		if err != nil {
			return jwtData, err
		}

		c.Header("authToken", newJwtToken)
		return jwtData, err
	}

	jwtData, err = jwt.DecodeBearer(jwtToken)
	if err != nil {
		jwtData, newJwtToken, err := CreateNewTokenWithRefreshToken(c, db)
		if err != nil {
			return jwtData, err
		}
		c.Header("authToken", newJwtToken)
		return jwtData, err
	}
	return jwtData, err
}

func CreateNewTokenWithRefreshToken(c *gin.Context, db *gorm.DB) (jwt.JWTPayload, string, error) {
	refreshToken, err := GetRefreshTokenFromHeader(c)
	var jwtData jwt.JWTPayload
	if err != nil {
		return jwtData, "", err
	}
	refreshTokenBody, err := jwt.VerifyRefreshToken(refreshToken, db)
	if err != nil {
		return jwtData, "", err
	}

	userData := jwt.JWTUser{
		UserId:   refreshTokenBody.UserId,
		Username: refreshTokenBody.Username,
	}

	jwtToken, err := jwt.CreateToken(userData)
	if err != nil {
		return jwtData, jwtToken, err
	}

	jwtData, err = jwt.DecodeBearer(jwtToken)
	return jwtData, jwtToken, err
}
