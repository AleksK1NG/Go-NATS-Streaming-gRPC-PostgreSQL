package nats

import "time"

const (
	ackWait     = 60 * time.Second
	durableName = "microservice-dur"
	maxInflight = 25

	createEmailWorkers = 6
	sendEmailWorkers   = 6

	createEmailSubject = "mail:create"
	sendEmailSubject   = "mail:send"
	emailGroupName     = "email_service"

	workerTimeout = 15 * time.Second
)
