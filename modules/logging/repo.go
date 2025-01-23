package logging

import "gorm.io/gorm"

func CreateLogEntryDB(entry NewLogEntry, db *gorm.DB) (string, error) {
	tx := db.Table("LogEntry").Create(&entry)
	//TODO: Figure out how to get the id of the log entry

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
}
