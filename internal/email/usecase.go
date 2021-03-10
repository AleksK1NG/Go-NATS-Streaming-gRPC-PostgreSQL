package email

import (
	"context"

	"github.com/AleksK1NG/nats-streaming/internal/models"
	"github.com/AleksK1NG/nats-streaming/pkg/utils"
	"github.com/satori/go.uuid"
)

// UseCase Email usecase interface
type UseCase interface {
	Create(ctx context.Context, email *models.Email) error
	PublishCreate(ctx context.Context, email *models.Email) error
	GetByID(ctx context.Context, emailID uuid.UUID) (*models.Email, error)
	Search(ctx context.Context, search string, pagination *utils.Pagination) (*models.EmailsList, error)
	SendEmail(ctx context.Context, email *models.Email) error
}
