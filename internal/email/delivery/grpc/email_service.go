package grpc

import (
	"context"

	"github.com/AleksK1NG/nats-streaming/internal/email"
	"github.com/AleksK1NG/nats-streaming/internal/models"
	grpcErrors "github.com/AleksK1NG/nats-streaming/pkg/grpc_errors"
	"github.com/AleksK1NG/nats-streaming/pkg/logger"
	"github.com/AleksK1NG/nats-streaming/pkg/utils"
	"github.com/AleksK1NG/nats-streaming/proto/email"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	uuid "github.com/satori/go.uuid"
)

type emailGRPCService struct {
	emailUC   email.UseCase
	log       logger.Logger
	validator *validator.Validate
}

// NewEmailGRPCService email gRPC service constructor
func NewEmailGRPCService(emailUC email.UseCase, log logger.Logger, validator *validator.Validate) *emailGRPCService {
	return &emailGRPCService{emailUC: emailUC, log: log, validator: validator}
}

// Create create email
func (e *emailGRPCService) Create(ctx context.Context, req *emailService.CreateReq) (*emailService.CreateRes, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productService.Create")
	defer span.Finish()
	createRequests.Inc()

	m := &models.Email{
		From:    req.GetFrom(),
		To:      req.GetTo(),
		Subject: req.GetSubject(),
		Message: req.GetMessage(),
	}

	if err := e.validator.StructCtx(ctx, m); err != nil {
		errorRequests.Inc()
		e.log.Errorf("validator.StructCtx: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	created, err := e.emailUC.Create(ctx, m)
	if err != nil {
		errorRequests.Inc()
		e.log.Errorf("validator.StructCtx: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	successRequests.Inc()
	return &emailService.CreateRes{Email: created.ToProto()}, nil
}

// GetByID find single email by id
func (e *emailGRPCService) GetByID(ctx context.Context, req *emailService.GetByIDReq) (*emailService.GetByIDRes, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productService.GetByID")
	defer span.Finish()
	getByIdRequests.Inc()

	emailUUID, err := uuid.FromString(req.GetEmailID())
	if err != nil {
		errorRequests.Inc()
		e.log.Errorf("uuid.FromString: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	m, err := e.emailUC.GetByID(ctx, emailUUID)
	if err != nil {
		errorRequests.Inc()
		e.log.Errorf("emailUC.GetByID: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	successRequests.Inc()
	return &emailService.GetByIDRes{Email: m.ToProto()}, nil
}

// Search find email by search text
func (e *emailGRPCService) Search(ctx context.Context, req *emailService.SearchReq) (*emailService.SearchRes, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productService.Search")
	defer span.Finish()
	searchRequests.Inc()

	res, err := e.emailUC.Search(ctx, req.GetSearch(), utils.NewPaginationQuery(int(req.GetSize()), int(req.GetPage())))
	if err != nil {
		errorRequests.Inc()
		e.log.Errorf("emailUC.GetByID: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	return &emailService.SearchRes{
		TotalCount: res.TotalCount,
		TotalPages: res.TotalPages,
		Page:       res.Page,
		Size:       res.Size,
		HasMore:    res.HasMore,
		Emails:     res.ToProto(),
	}, nil
}
