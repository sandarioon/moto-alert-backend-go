package email

type Service interface {
	SendTransactionalEmail(senderName string, senderEmail string, recipientEmail string, subject string, htmlContent string, params map[string]string) error
}
