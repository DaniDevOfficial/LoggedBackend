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
