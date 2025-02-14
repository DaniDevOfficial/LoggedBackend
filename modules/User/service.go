package User

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"loggedin/utility/hashing"
	"loggedin/utility/jwt"
	"net/http"
)

func Login(c *gin.Context, db *gorm.DB) {
	var loginData LoginRequest
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: "Invalid Request Body"})
		return
	}

	userData, err := GetUserInformationByUsername(loginData.Username, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{Message: "Internal server Error"})
		return
	}

	if userData.Username != loginData.Username {
		c.JSON(http.StatusInternalServerError, Error{Message: "Wrong username or password"})
		return
	}

	if !hashing.CheckHashedString(userData.Password, loginData.Password) {
		c.JSON(http.StatusBadRequest, Error{Message: "Wrong username or password"})
		return
	}
	jwtUser := jwt.JWTUser{
		UserId:   userData.Id,
		Username: userData.Username,
	}
	if userData.IsClaimed {

		refreshToken, err := jwt.CreateRefreshToken(jwtUser, loginData.IsTimeBased, db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Error{Message: "Internal server Error"})
			return
		}
		token, err := jwt.CreateToken(jwtUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Error{Message: "Internal server Error"})
			return
		}

		c.Writer.Header().Set("Authorization", token)
		c.Writer.Header().Set("RefreshToken", refreshToken)
	} else {
		token, err := jwt.CreateClaimToken(jwtUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Error{Message: "Internal server Error"})
			return
		}
		c.Writer.Header().Set("ClaimToken", token)
	}
	c.JSON(http.StatusOK, LoginResponse{IsClaimed: userData.IsClaimed})
}
