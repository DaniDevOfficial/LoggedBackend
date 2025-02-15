package User

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUseroutes(router *gin.Engine, db *gorm.DB) {
	registerAuthRoutes(router, db)
}

func registerAuthRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/auth/login", func(c *gin.Context) {
		Login(c, db)
	})

	router.POST("/auth/claim", func(c *gin.Context) {
		Claim(c, db)
	})

	router.PUT("/auth/password", func(c *gin.Context) {

	})
}
