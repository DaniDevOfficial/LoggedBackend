package main

import (
	"fmt"
	"loggedin/modules/logging"
)

func main() {
	jsonInput := `{"name": "John", "age": 30, "isAdmin": true, "password": 123}`

	tmp := logging.EncodePersonalInformation(jsonInput)
	fmt.Println(tmp)
	/*dbConnection := db.InitDB()
	router := gin.Default()
	validator.InitCustomValidators()

	logging.RegisterLoggingRoutes(router, dbConnection)

	err := router.Run("localhost:8000")
	if err != nil {
		panic("Startup went wrong")
	}*/
}
