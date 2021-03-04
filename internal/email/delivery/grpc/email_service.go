package grpc

import (
	"context"

	"github.com/AleksK1NG/nats-streaming/internal/email"
	"github.com/AleksK1NG/nats-streaming/pkg/logger"
	"github.com/AleksK1NG/nats-streaming/proto/email"
	"github.com/go-playground/validator/v10"
)

type emailGRPCService struct {
	emailUC   email.UseCase
	log       logger.Logger
	validator *validator.Validate
}

func NewEmailGRPCService(emailUC email.UseCase, log logger.Logger, validator *validator.Validate) *emailGRPCService {
	return &emailGRPCService{emailUC: emailUC, log: log, validator: validator}
}

func (e *emailGRPCService) Create(ctx context.Context, req *emailService.CreateReq) (*emailService.CreateRes, error) {
	panic("implement me")
}

func (e *emailGRPCService) GetByID(ctx context.Context, req *emailService.GetByIDReq) (*emailService.GetByIDRes, error) {
	panic("implement me")
}

func (e *emailGRPCService) Search(ctx context.Context, req *emailService.SearchReq) (*emailService.SearchRes, error) {
	panic("implement me")
}
