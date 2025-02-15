package User

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"loggedin/utility/auth"
	"loggedin/utility/hashing"
	"loggedin/utility/jwt"
	"loggedin/utility/validation"
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
		if errors.Is(err, NotFoundError) {
			c.JSON(http.StatusBadRequest, Error{Message: "Wrong username or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, Error{Message: "Internal server Error"})
		return
	}

	if userData.Username != loginData.Username {
		c.JSON(http.StatusBadRequest, Error{Message: "Wrong username or password"})
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

func Claim(c *gin.Context, db *gorm.DB) {
	var claimData ClaimRequest
	if err := c.ShouldBindJSON(&claimData); err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: "Invalid Request Body"})
		return
	}

	c.Set("isClaimRequest", true)
	jwtToken, err := auth.GetJWTPayloadFromHeader(c, db)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Error{Message: "IDK"})
		return
	}

	isValid, err := validation.IsValidPassword(claimData.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
		return
	}
	if !isValid {
		c.JSON(http.StatusBadRequest, Error{Message: "IDK"})
		return
	}

	userData, err := GetUserInformationById(jwtToken.UserId, db)
	if err != nil {
		if errors.Is(err, NotFoundError) {
			c.JSON(http.StatusBadRequest, Error{Message: "User Not Found"})
			return
		}
		c.JSON(http.StatusInternalServerError, Error{Message: "Internal server Error"})
		return
	}

	if userData.IsClaimed {
		c.JSON(http.StatusBadRequest, Error{Message: "User Claim Error"})
		return
	}
	jwtUser := jwt.JWTUser{
		UserId:   userData.Id,
		Username: userData.Username,
	}

	refreshToken, err := jwt.CreateRefreshToken(jwtUser, false, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{Message: "Internal server Error"})
		return
	}

	token, err := jwt.CreateToken(jwtUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{Message: "Internal server Error"})
		return
	}
	err = MarkUserAsClaimed(userData.Id, db)
	if err != nil {
		if errors.Is(err, NotFoundError) {
			c.JSON(http.StatusBadRequest, Error{Message: "User Not Found"})
			return
		}
		c.JSON(http.StatusInternalServerError, Error{Message: "Internal server Error"})
		return
	}
	c.Writer.Header().Set("Authorization", token)
	c.Writer.Header().Set("RefreshToken", refreshToken)
	c.JSON(http.StatusOK, claimData)
}
