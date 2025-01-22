package logging

import "database/sql"

func CreateLogEntryDB(entry NewLogEntry, db *sql.DB) (string, error) {
	query := `
INSERT INTO LogEntry
    (severity, message, request, user_id, request_url, response, lifetime, request_key, date_time) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING
	id
`
	row := db.QueryRow(query, entry.Severity, entry.Message, entry.Request, entry.UserId, entry.RequestUrl, entry.Response, entry.LifeTime, entry.RequestKey, entry.DateTime)

	var id string
	err := row.Scan(&id)
	return id, err
}

func GetFilteredLogEntriesFromDB(db *gorm.DB, filters FilterLogEntryRequest) () {
	query := db.table('LogEntry')

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
		query = query.Where("")
	}
}
