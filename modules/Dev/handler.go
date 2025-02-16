package Dev

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"loggedin/utility/auth"
	"loggedin/utility/hashing"
	"net/http"
)

type HashingRequest struct {
	Password string `json:"password"`
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

}
