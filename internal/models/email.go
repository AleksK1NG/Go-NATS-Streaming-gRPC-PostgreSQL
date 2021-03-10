package models

import (
	"time"

	emailService "github.com/AleksK1NG/nats-streaming/proto/email"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Email model
type Email struct {
	EmailID   uuid.UUID `json:"emailID"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

// EmailsList emails list response with pagination
type EmailsList struct {
	TotalCount int64    `json:"totalCount"`
	TotalPages int64    `json:"totalPages"`
	Page       int64    `json:"page"`
	Size       int64    `json:"size"`
	HasMore    bool     `json:"hasMore"`
	Emails     []*Email `json:"emails"`
}

// ToProto convert email to proto
func (e *Email) ToProto() *emailService.Email {
	return &emailService.Email{
		EmailID:   e.EmailID.String(),
		From:      e.From,
		To:        e.To,
		Subject:   e.Subject,
		Message:   e.Subject,
		CreatedAt: timestamppb.New(e.CreatedAt),
	}
}

func (l *EmailsList) ToProto() []*emailService.Email {
	mails := make([]*emailService.Email, 0, len(l.Emails))
	for _, e := range l.Emails {
		mails = append(mails, e.ToProto())
	}
	return mails
}

// MailData for send email
type MailData struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	// Content string `json:"content"`
	Content string `json:"content"`
}
