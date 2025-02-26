package User

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.Engine, db *gorm.DB) {
	registerAuthRoutes(router, db)
}

func registerAuthRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/auth/login", func(c *gin.Context) {
		Login(c, db)
	})

	router.POST("/auth/claim", func(c *gin.Context) {
		Claim(c, db)
	})

	router.POST("/auth/account", func(c *gin.Context) {
		CreateNewClaimAccount(c, db)
	})

	router.DELETE("/auth/account", func(c *gin.Context) {
		DeleteAccount(c, db)
	})

	router.POST("/auth/roles/admin", func(c *gin.Context) {
		AddAdminRoleToUser(c, db)
	})
	router.DELETE("/auth/roles/admin", func(c *gin.Context) {
		RemoveAdminRoleFromUser(c, db)
	})

	router.GET("/auth/check", func(c *gin.Context) {
		CheckAuth(c, db)
	})
}
