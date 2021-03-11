package usecase

import (
	"context"
	"encoding/json"

	"github.com/AleksK1NG/nats-streaming/internal/email"
	"github.com/AleksK1NG/nats-streaming/internal/email/delivery/nats"
	"github.com/AleksK1NG/nats-streaming/internal/models"
	smtpClient "github.com/AleksK1NG/nats-streaming/pkg/email"
	"github.com/AleksK1NG/nats-streaming/pkg/logger"
	"github.com/AleksK1NG/nats-streaming/pkg/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

const (
	createEmailSubject = "mail:create"
	sendEmailSubject   = "mail:send"
)

type emailUseCase struct {
	log         logger.Logger
	emailPGRepo email.PGRepository
	publisher   nats.Publisher
	smtpClient  smtpClient.SMTPClient
}

// NewEmailUseCase email usecase constructor
func NewEmailUseCase(log logger.Logger, emailPGRepo email.PGRepository, publisher nats.Publisher, smtpClient smtpClient.SMTPClient) *emailUseCase {
	return &emailUseCase{log: log, emailPGRepo: emailPGRepo, publisher: publisher, smtpClient: smtpClient}
}

// Create create new email saves in db
func (e *emailUseCase) Create(ctx context.Context, email *models.Email) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "emailUseCase.Create")
	defer span.Finish()

	created, err := e.emailPGRepo.Create(ctx, email)
	if err != nil {
		return errors.Wrap(err, "emailPGRepo.Create")
	}

	mailBytes, err := json.Marshal(created)
	if err != nil {
		return errors.Wrap(err, "json.Marshal")
	}

	return e.publisher.Publish(sendEmailSubject, mailBytes)
}

// GetByID fnd email by id
func (e *emailUseCase) GetByID(ctx context.Context, emailID uuid.UUID) (*models.Email, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "emailUseCase.GetByID")
	defer span.Finish()
	return e.emailPGRepo.GetByID(ctx, emailID)
}

// PublishCreate publish create email event to message broker
func (e *emailUseCase) PublishCreate(ctx context.Context, email *models.Email) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "emailUseCase.PublishCreate")
	defer span.Finish()

	mailBytes, err := json.Marshal(email)
	if err != nil {
		return errors.Wrap(err, "json.Marshal")
	}

	return e.publisher.Publish(createEmailSubject, mailBytes)
}

// Search search email in db
func (e *emailUseCase) Search(ctx context.Context, search string, pagination *utils.Pagination) (*models.EmailsList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "emailUseCase.Search")
	defer span.Finish()
	return e.emailPGRepo.Search(ctx, search, pagination)
}

// SendEmail send email
func (e *emailUseCase) SendEmail(ctx context.Context, email *models.Email) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "emailUseCase.SendEmail")
	defer span.Finish()

	return e.smtpClient.SendMail(&models.MailData{
		To:      email.To,
		From:    email.From,
		Subject: email.Subject,
		Content: email.Message,
	})
}
