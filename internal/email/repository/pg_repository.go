package repository

import (
	"context"

	"github.com/AleksK1NG/nats-streaming/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/satori/go.uuid"
)

type emailPGRepository struct {
	db *pgxpool.Pool
}

func NewEmailPGRepository(db *pgxpool.Pool) *emailPGRepository {
	return &emailPGRepository{db: db}
}

func (e *emailPGRepository) Create(ctx context.Context, email *models.Email) (*models.Email, error) {
	panic("implement me")
}

func (e *emailPGRepository) GetByID(ctx context.Context, emailID uuid.UUID) (*models.Email, error) {
	panic("implement me")
}
