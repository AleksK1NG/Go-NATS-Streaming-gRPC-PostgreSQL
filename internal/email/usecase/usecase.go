package usecase

import (
	"context"
	"encoding/json"

	"github.com/AleksK1NG/nats-streaming/internal/email"
	"github.com/AleksK1NG/nats-streaming/internal/email/delivery/nats"
	"github.com/AleksK1NG/nats-streaming/internal/models"
	"github.com/AleksK1NG/nats-streaming/pkg/logger"
	smtpClient "github.com/AleksK1NG/nats-streaming/pkg/smtp"
	"github.com/AleksK1NG/nats-streaming/pkg/utils"
	"github.com/go-redis/redis/v8"
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
	redisRepo   email.RedisRepository
}

// NewEmailUseCase email usecase constructor
func NewEmailUseCase(log logger.Logger, emailPGRepo email.PGRepository, publisher nats.Publisher, smtpClient smtpClient.SMTPClient, redisRepo email.RedisRepository) *emailUseCase {
	return &emailUseCase{log: log, emailPGRepo: emailPGRepo, publisher: publisher, smtpClient: smtpClient, redisRepo: redisRepo}
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

	cached, err := e.redisRepo.GetEmailByID(ctx, emailID)
	if err != nil && err != redis.Nil {
		e.log.Errorf("redisRepo.GetEmailByID: %v", err)
	}
	if cached != nil {
		return cached, nil
	}

	mail, err := e.emailPGRepo.GetByID(ctx, emailID)
	if err != nil {
		return nil, errors.Wrap(err, "emailPGRepo.GetByID")
	}

	if err := e.redisRepo.SetEmail(ctx, mail); err != nil {
		e.log.Errorf("redisRepo.SetEmail: %v", err)
	}

	return mail, nil
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

// SendEmail send email using smtp client
func (e *emailUseCase) SendEmail(ctx context.Context, email *models.Email) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "emailUseCase.SendEmail")
	defer span.Finish()

	if err := e.smtpClient.SendMail(&models.MailData{
		To:      email.To,
		From:    email.From,
		Subject: email.Subject,
		Content: email.Message,
	}); err != nil {
		return errors.Wrap(err, "SendMail")
	}

	return nil
}
