package Dev

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"loggedin/utility/auth"
)

func RegisterDevRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/dev/jwtTest", func(c *gin.Context) {
		//This is for testing the refresh token and authentication token and how a new one gets generated
		payload, err := auth.GetJWTPayloadFromHeader(c, db)
		fmt.Println(payload)
		fmt.Println(err)
	})
}
