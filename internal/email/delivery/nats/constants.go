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

	deadLetterQueueSubject = "mail:errors"
	maxRedeliveryCount     = 3
)
