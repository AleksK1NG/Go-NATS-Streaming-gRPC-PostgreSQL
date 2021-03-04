package email

import (
	"context"

	"github.com/AleksK1NG/nats-streaming/internal/models"
	uuid "github.com/satori/go.uuid"
)

// PGRepository Email postgresql repository
type PGRepository interface {
	Create(ctx context.Context, email *models.Email) (*models.Email, error)
	GetByID(ctx context.Context, emailID uuid.UUID) (*models.Email, error)
}
