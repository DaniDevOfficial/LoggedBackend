package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"loggedin/modules/Dev"
	"loggedin/modules/logging"
	"loggedin/utility/db"
	"loggedin/utility/validator"
	"net/http"
)

func main() {
	dbConnection := db.InitDB()
	router := gin.Default()
	router.Use(corsMiddleware())
	validator.InitCustomValidators()

	logging.RegisterLoggingRoutes(router, dbConnection)
	Dev.RegisterDevRoutes(router, dbConnection)
	err := router.Run("localhost:8000")
	if err != nil {
		panic("Startup went wrong")
	}
}

// corsMiddleware sets the CORS headers to allow all origins.
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, RefreshToken")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			log.Println("Options request handled")
			return
		}

		log.Println("New Request Started")
		log.Printf("Method: %s, Path: %s\n", c.Request.Method, c.Request.URL.Path)

		c.Next()
	}
}
