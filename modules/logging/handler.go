package logging

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func RegisterLoggingRoutes(router *gin.Engine, db *sql.DB) {
	registerBasicLoggingRoutes(router, db)
}

func registerBasicLoggingRoutes(router *gin.Engine, db *sql.DB) {
	router.POST("/logEntry", func(c *gin.Context) {
		CreateLogEntry(c, db)
	})
	router.GET("/logEntry", func(c *gin.Context) {
		GetFilteredLogEntriesWithLimit(c, db)
	})
}
