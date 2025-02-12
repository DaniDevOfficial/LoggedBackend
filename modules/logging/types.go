package logging

type NewLogEntry struct {
	ID         string `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Auto-generate UUID
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
	LogEntryId       string `form:"logEntryId"`
	SeverityFilter   string `form:"severityFilter"`
	MessageFilter    string `form:"messageFilter"`
	RequestFilter    string `form:"requestFilter"`
	UserIdFilter     string `form:"userIdFilter"`
	RequestUrlFilter string `form:"requestUrlFilter"`
	ResponseFilter   string `form:"responseFilter"`
	LifeTimeFilter   string `form:"lifeTimeFilter"`
	RequestKeyFilter string `form:"requestKeyFilter"`
	StartDateFilter  string `form:"startDateFilter"`
	EndDateFilter    string `form:"endDateFilter"`
	Limit            int    `form:"limit"`
	Page             int    `form:"page"`
	Ordering         string `form:"ordering"`
}

type Error struct {
	Message string `json:"message"`
}

type IdResponse struct {
	Id string `json:"id"`
}
