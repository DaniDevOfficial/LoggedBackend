package logging

import (
	"gorm.io/gorm"
)

func CreateLogEntryDB(entry NewLogEntry, db *gorm.DB) (string, error) {
	result := db.Table("logs").Create(&entry)

	if result.Error != nil {
		return "", result.Error
	}

	return entry.ID, nil
}

func GetFilteredLogEntriesFromDB(db *gorm.DB, filters FilterLogEntryRequest) {
	query := db.Table("LogEntry")

	if filters.LogEntryId != "" {
		query = query.Where("id = ?", filters.LogEntryId)
	}

	if filters.SeverityFilter != "" {
		query = query.Where("severity = ?", filters.SeverityFilter)
	}
	if filters.MessageFilter != "" {
		query = query.Where("message ILIKE ?", "%"+filters.MessageFilter+"%")
	}
	if filters.RequestFilter != "" {
		query = query.Where("request ILIKE ?", "%"+filters.RequestFilter+"%")
	}
	if filters.UserIdFilter != "" {
		query = query.Where("user_id = ?", filters.UserIdFilter)
	}
	if filters.RequestUrlFilter != "" {
		query = query.Where("request_url ILIKE ?", "%"+filters.RequestUrlFilter+"%")
	}
	if filters.ResponseFilter != "" {
		query = query.Where("response ILIKE ?", "%"+filters.ResponseFilter+"%")
	}
	if filters.RequestKeyFilter != "" {
		query = query.Where("request_key ILIKE ?", "%"+filters.RequestKeyFilter+"%")
	}
	if filters.StartDateFilter != "" {
		query = query.Where("dateTime >= ?", filters.StartDateFilter)
	}
	if filters.EndDateFilter != "" {
		query = query.Where("dateTime <= ?", filters.EndDateFilter)
	}
	//TODO: figure out gorm
	if filters.Ordering == "asc" {
		query = query.Order("date_time asc")
	} else {
		query = query.Order("date_time desc")

	}
	//Scan row for each entry save it and then return the array of entries
	return
}
