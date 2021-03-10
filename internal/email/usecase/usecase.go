package usecase

import (
	"context"
	"encoding/json"

	"github.com/AleksK1NG/nats-streaming/internal/email"
	"github.com/AleksK1NG/nats-streaming/internal/email/delivery/nats"
	"github.com/AleksK1NG/nats-streaming/internal/models"
	"github.com/AleksK1NG/nats-streaming/pkg/logger"
	"github.com/AleksK1NG/nats-streaming/pkg/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

const (
	createEmailSubject = "mail:create"
)

type emailUseCase struct {
	log         logger.Logger
	emailPGRepo email.PGRepository
	publisher   nats.Publisher
}

// NewEmailUseCase email usecase constructor
func NewEmailUseCase(log logger.Logger, emailPGRepo email.PGRepository, publisher nats.Publisher) *emailUseCase {
	return &emailUseCase{log: log, emailPGRepo: emailPGRepo, publisher: publisher}
}

// Create create new email saves in db
func (e *emailUseCase) Create(ctx context.Context, email *models.Email) (*models.Email, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "emailUseCase.Create")
	defer span.Finish()
	return e.emailPGRepo.Create(ctx, email)
}

// GetByID fnd email by id
func (e *emailUseCase) GetByID(ctx context.Context, emailID uuid.UUID) (*models.Email, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "emailUseCase.GetByID")
	defer span.Finish()
	return e.emailPGRepo.GetByID(ctx, emailID)
}

// PublishCreate publish create email event to message broker
func (e *emailUseCase) PublishCreate(ctx context.Context, email *models.Email) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "emailUseCase.PublishCreate")
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
