package Dev

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"loggedin/modules/User"
	"loggedin/utility/auth"
	"loggedin/utility/hashing"
	"net/http"
)

type HashingRequest struct {
	Password string `json:"password"`
}
type ComparingRequest struct {
	Password string `json:"password"`
	Hash     string `json:"hash"`
}

func RegisterDevRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/dev/jwtTest", func(c *gin.Context) {
		//This is for testing the refresh token and authentication token and how a new one gets generated
		payload, err := auth.GetJWTPayloadFromHeader(c, db)
		fmt.Println(payload)
		fmt.Println(err)
	})

	router.POST("/dev/hashing", func(c *gin.Context) {
		var request HashingRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		password, err := hashing.HashPassword(request.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"hashedPassword": password})
	})
	router.POST("/dev/compare", func(c *gin.Context) {
		var request ComparingRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"is the same": hashing.CheckHashedString(request.Hash, request.Password)})
	})

	router.GET("gimme", func(c *gin.Context) {

		data, err := auth.GetJWTPayloadFromHeader(c, db)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "noooo")
			return
		}

		err = User.AddUserAdmin(data.UserId, db)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}

	})

}
