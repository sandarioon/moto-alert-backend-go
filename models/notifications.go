package models

type EmailTemplate struct {
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PushNotificationTemplate struct {
	Id         int    `json:"id"`
	Type       string `json:"type"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Recipients string `json:"recipients"`
}

type EmailType string

const (
	SEND_CODE                EmailType = "SEND_CODE"
	SEND_NEW_PASSWORD        EmailType = "SEND_NEW_PASSWORD"
	SEND_DELETE_ACCOUNT_CODE EmailType = "SEND_DELETE_ACCOUNT_CODE"
)
