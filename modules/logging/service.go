package logging

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateLogEntry(c *gin.Context, db *sql.DB) {

	var newLogEntry NewLogEntry

	if err := c.ShouldBindJSON(&newLogEntry); err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: "Invalid Request Body"})
		return
	}

	//TODO: Auth

	//TODO: Check for passwords etc inside of the log entry and hide them with ***

	//TODO: save entry in database and response the created Id
	id, err := CreateLogEntryDB(newLogEntry, db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{Message: "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, IdResponse{Id: id})
}
