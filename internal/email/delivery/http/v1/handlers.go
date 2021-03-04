package v1

import (
	"github.com/AleksK1NG/nats-streaming/internal/email"
	"github.com/AleksK1NG/nats-streaming/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type emailHandlers struct {
	group     *echo.Group
	emailUC   email.UseCase
	log       logger.Logger
	validator *validator.Validate
}

func NewEmailHandlers(group *echo.Group, emailUC email.UseCase, log logger.Logger, validator *validator.Validate) *emailHandlers {
	return &emailHandlers{group: group, emailUC: emailUC, log: log, validator: validator}
}

func (h *emailHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(200, "ok")
	}
}

func (h *emailHandlers) GetByID() echo.HandlerFunc {
	panic("implement me")
}

func (h *emailHandlers) Search() echo.HandlerFunc {
	panic("implement me")
}
