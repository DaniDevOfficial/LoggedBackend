package main

import (
	"github.com/gin-gonic/gin"
	"loggedin/modules/logging"
	"loggedin/utility/db"
)

func main() {
	dbConnection := db.InitDB()
	router := gin.Default()

	logging.RegisterLoggingRoutes(router, dbConnection)

	err := router.Run("localhost:8000")
	if err != nil {
		panic(123)
	}
}
