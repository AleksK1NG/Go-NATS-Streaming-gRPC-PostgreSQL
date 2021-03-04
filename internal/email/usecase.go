package email

import (
	"context"

	"github.com/AleksK1NG/nats-streaming/internal/models"
	"github.com/satori/go.uuid"
)

// UseCase Email usecase interface
type UseCase interface {
	Create(ctx context.Context, email *models.Email) (*models.Email, error)
	GetByID(ctx context.Context, emailID uuid.UUID) (*models.Email, error)
}
