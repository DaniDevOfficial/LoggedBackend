package main

import (
	"github.com/gin-gonic/gin"
	"loggedin/modules/logging"
	"loggedin/utility/db"
	"loggedin/utility/validator"
)

func main() {
	dbConnection := db.InitDB()
	router := gin.Default()
	validator.InitCustomValidators()

	logging.RegisterLoggingRoutes(router, dbConnection)

	err := router.Run("localhost:8000")
	if err != nil {
		panic("Startup went wrong")
	}
}
