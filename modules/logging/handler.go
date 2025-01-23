package logging

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterLoggingRoutes(router *gin.Engine, db *gorm.DB) {
	registerBasicLoggingRoutes(router, db)
}

func registerBasicLoggingRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/logEntry", func(c *gin.Context) {
		CreateLogEntry(c, db)
	})
	router.GET("/logEntry", func(c *gin.Context) {
		GetFilteredLogEntriesWithLimit(c, db)
	})
}
