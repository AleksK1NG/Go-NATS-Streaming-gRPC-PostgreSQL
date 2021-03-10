package nats

const (
	createEmailWorkers = 6
	sendEmailWorkers   = 6

	createEmailSubject = "mail:create"
	sendEmailSubject   = "mail:send"
	emailGroupName     = "email_service"
)
