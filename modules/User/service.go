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
	"time"
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
		log.Println("Wrong password")
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
		c.Writer.Header().Set("Authorization", token)
	}
	c.JSON(http.StatusOK, LoginResponse{IsClaimed: userData.IsClaimed})
}

type ClaimUser struct {
	IsClaimed bool   `json:"is_claimed"`
	Password  string `json:"password"`
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
		c.JSON(http.StatusUnauthorized, Error{Message: "Please try again"})
		return
	}
	if !jwtToken.IsClaimToken {
		c.JSON(http.StatusBadRequest, Error{Message: "IDK"})
		return
	}
	isValid, err := validation.IsValidPassword(claimData.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
		return
	}
	if !isValid {
		c.JSON(http.StatusBadRequest, Error{Message: "how???"})
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

	hashedPassword, err := hashing.HashPassword(claimData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{Message: "Internal server Error"})
		return
	}

	claimUser := ClaimUser{
		IsClaimed: true,
		Password:  hashedPassword,
	}
	err = MarkUserAsClaimed(userData.Id, claimUser, db)
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

func CheckAuth(c *gin.Context, db *gorm.DB) {
	time.Sleep(2 * time.Second)

	_, err := auth.GetJWTPayloadFromHeader(c, db)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, Error{Message: "Unauthorized to perform action please login to continue"})
		return
	}

	c.JSON(http.StatusOK, Success{Message: "Authenticated"})
}

type NewAccountRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateNewClaimAccount(c *gin.Context, db *gorm.DB) {
	var newAccountData NewAccountRequest
	if err := c.ShouldBindJSON(&newAccountData); err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: "Invalid Request Body"})
		return
	}

	jwtData, err := auth.GetJWTPayloadFromHeader(c, db)
	if err != nil {
		c.JSON(http.StatusUnauthorized, Error{Message: "Unauthorized"})
		return
	}

	if !IsUserAdmin(jwtData.UserId, db) {
		c.JSON(http.StatusForbidden, Error{Message: "You are not allowed to perform this action"})
		return
	}
	if !UsernameAlreadyInUse(newAccountData.Username, db) {
		c.JSON(http.StatusBadRequest, Error{Message: "Username Already In Use"})
		return
	}

	_, err = validation.IsValidPassword(newAccountData.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
		return
	}
	newAccountData.Password, err = hashing.HashPassword(newAccountData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{Message: "Internal server Error"})
		return
	}

	userId, err := CreateNewUser(newAccountData, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{Message: "Internal server Error"})
		return
	}

	c.JSON(http.StatusOK, SuccessWithId{Message: "New Account Created", Id: userId})
}

type DeleteRequest struct {
	Id string `form:"id" binding:"required,uuid"`
}

func DeleteAccount(c *gin.Context, db *gorm.DB) {
	var deleteRequest DeleteRequest
	if err := c.ShouldBindQuery(&deleteRequest); err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: "Invalid Request Body"})
	}

	jwtData, err := auth.GetJWTPayloadFromHeader(c, db)
	if err != nil {
		c.JSON(http.StatusUnauthorized, Error{Message: "Unauthorized"})
		return
	}

	if jwtData.UserId != deleteRequest.Id || !IsUserAdmin(jwtData.UserId, db) {
		c.JSON(http.StatusForbidden, Error{Message: "You are not allowed to perform this action"})
		return
	}
	err = DeleteAccountInDB(deleteRequest.Id, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{Message: "Internal server Error"})
		return
	}

	c.JSON(http.StatusOK, Success{Message: "Account Deleted"})
}

type Success struct {
	Message string `json:"message"`
}
type SuccessWithId struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}
