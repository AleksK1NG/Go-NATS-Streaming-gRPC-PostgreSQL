package usecase

import (
	"context"

	"github.com/AleksK1NG/nats-streaming/internal/email"
	"github.com/AleksK1NG/nats-streaming/internal/models"
	"github.com/AleksK1NG/nats-streaming/pkg/logger"
	"github.com/satori/go.uuid"
)

type emailUseCase struct {
	log         logger.Logger
	emailPGRepo email.PGRepository
}

func NewEmailUseCase(log logger.Logger, emailPGRepo email.PGRepository) *emailUseCase {
	return &emailUseCase{log: log, emailPGRepo: emailPGRepo}
}

func (e *emailUseCase) Create(ctx context.Context, email *models.Email) (*models.Email, error) {
	panic("implement me")
}

func (e *emailUseCase) GetByID(ctx context.Context, emailID uuid.UUID) (*models.Email, error) {
	panic("implement me")
}
