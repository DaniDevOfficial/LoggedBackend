package logging

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

const HardLimitEntries = 500

const DefaultLimitEntries = 50

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

func GetFilteredLogEntriesWithLimit(c *gin.Context, db *sql.DB) {
	var filters FilterLogEntryRequest
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: "Invalid Request Body"})
		return
	}

	if filters.Limit == 0 {
		filters.Limit = DefaultLimitEntries
	}
	if filters.Limit > HardLimitEntries {
		filters.Limit = HardLimitEntries
	}

}
