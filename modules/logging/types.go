package logging

type NewLogEntry struct {
	Severity   string `json:"severity" binding:"required"`
	Message    string `json:"message" binding:"required"`
	Request    string `json:"request" binding:"required"`
	UserId     string `json:"userId" binding:"required"`
	RequestUrl string `json:"requestUrl" binding:"required"`
	Response   string `json:"response" binding:"required"`
	LifeTime   string `json:"lifeTime" binding:"required"`
	RequestKey string `json:"requestKey" binding:"required"`
	DateTime   string `json:"dateTime" binding:"required,dateTime"`
}

type FilterLogEntryRequest struct {
	LogEntryId       string `json:"logEntryId"`
	SeverityFilter   string `json:"severityFilter"`
	MessageFilter    string `json:"messageFilter"`
	RequestFilter    string `json:"requestFilter"`
	UserIdFilter     string `json:"userIdFilter"`
	RequestUrlFilter string `json:"requestUrlFilter"`
	ResponseFilter   string `json:"responseFilter"`
	RequestKeyFilter string `json:"requestKeyFilter"`
	StartDateFilter  string `json:"startDateFilter"`
	EndDateFilter    string `json:"endDateFilter"`
	Limit            int    `json:"limit"`
}

type Error struct {
	Message string `json:"message"`
}

type IdResponse struct {
	Id string `json:"id"`
}
