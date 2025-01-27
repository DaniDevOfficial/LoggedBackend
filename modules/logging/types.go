package logging

type NewLogEntry struct {
	ID         string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Auto-generate UUID
	Severity   string `json:"severity" binding:"required"`
	Message    string `json:"message"`
	Request    string `json:"request"`
	UserId     string `json:"userId"`
	RequestUrl string `json:"requestUrl"`
	Response   string `json:"response"`
	LifeTime   string `json:"lifeTime"`
	RequestKey string `json:"requestKey"`
	DateTime   string `json:"dateTime" binding:"dateTime"`
}

type EntryFromDB struct {
	Id         string `json:"id"`
	Severity   string `json:"severity"`
	Message    string `json:"message"`
	Request    string `json:"request"`
	UserId     string `json:"userId"`
	RequestUrl string `json:"requestUrl"`
	Response   string `json:"response"`
	LifeTime   string `json:"lifeTime"`
	RequestKey string `json:"requestKey"`
	DateTime   string `json:"dateTime"`
}

type Entry struct {
}

type FilterLogEntryRequest struct {
	LogEntryId       string `json:"logEntryId"`
	SeverityFilter   string `json:"severityFilter"`
	MessageFilter    string `json:"messageFilter"`
	RequestFilter    string `json:"requestFilter"`
	UserIdFilter     string `json:"userIdFilter"`
	RequestUrlFilter string `json:"requestUrlFilter"`
	ResponseFilter   string `json:"responseFilter"`
	LifeTimeFilter   string `json:"lifeTimeFilter"`
	RequestKeyFilter string `json:"requestKeyFilter"`
	StartDateFilter  string `json:"startDateFilter"`
	EndDateFilter    string `json:"endDateFilter"`
	Limit            int    `json:"limit"`
	Page             int    `json:"page"`
	Ordering         string `json:"ordering"`
}

type Error struct {
	Message string `json:"message"`
}

type IdResponse struct {
	Id string `json:"id"`
}
