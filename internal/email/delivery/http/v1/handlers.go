package v1

import (
	"net/http"
	"strconv"

	"github.com/AleksK1NG/nats-streaming/internal/email"
	"github.com/AleksK1NG/nats-streaming/internal/models"
	httpErrors "github.com/AleksK1NG/nats-streaming/pkg/http_errors"
	"github.com/AleksK1NG/nats-streaming/pkg/logger"
	"github.com/AleksK1NG/nats-streaming/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	uuid "github.com/satori/go.uuid"
)

type emailHandlers struct {
	group    *echo.Group
	emailUC  email.UseCase
	log      logger.Logger
	validate *validator.Validate
}

func NewEmailHandlers(group *echo.Group, emailUC email.UseCase, log logger.Logger, validate *validator.Validate) *emailHandlers {
	return &emailHandlers{group: group, emailUC: emailUC, log: log, validate: validate}
}

// Create Create
// @Tags Emails
// @Summary Create new email
// @Description Create new email and send it
// @Accept json
// @Produce json
// @Success 201 {object} models.Email
// @Router /email [post]
func (h *emailHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "emailHandlers.Create")
		defer span.Finish()
		createRequests.Inc()

		var mail models.Email
		if err := c.Bind(&mail); err != nil {
			errorRequests.Inc()
			h.log.Errorf("c.Bind: %v", err)
			return httpErrors.ErrorCtxResponse(c, err)
		}

		if err := h.validate.StructCtx(ctx, &mail); err != nil {
			errorRequests.Inc()
			h.log.Errorf("validate.StructCtx: %v", err)
			return httpErrors.ErrorCtxResponse(c, err)
		}

		if err := h.emailUC.PublishCreate(ctx, &mail); err != nil {
			errorRequests.Inc()
			h.log.Errorf("validate.StructCtx: %v", err)
			return httpErrors.ErrorCtxResponse(c, err)
		}

		successRequests.Inc()
		return c.NoContent(http.StatusCreated)
	}
}

// GetByID GetByID
// @Tags Emails
// @Summary Get email by id
// @Description Get email by email uuid
// @Accept json
// @Produce json
// @Param email_id path string true "email_id"
// @Success 200 {object} models.Email
// @Router /email/{email_id} [get]
func (h *emailHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "emailHandlers.GetByID")
		defer span.Finish()
		getByIdRequests.Inc()

		emailUUID, err := uuid.FromString(c.Param("email_id"))
		if err != nil {
			errorRequests.Inc()
			h.log.Errorf("uuid.FromString: %v", err)
			return httpErrors.ErrorCtxResponse(c, err)
		}

		m, err := h.emailUC.GetByID(ctx, emailUUID)
		if err != nil {
			errorRequests.Inc()
			h.log.Errorf("uuid.FromString: %v", err)
			return httpErrors.ErrorCtxResponse(c, err)
		}

		successRequests.Inc()
		return c.JSON(http.StatusOK, m)
	}
}

// Search Search emails
// @Tags Emails
// @Summary Search emails
// @Description Search email
// @Accept json
// @Produce json
// @Param search query string false "search text"
// @Param page query string false "page number"
// @Param size query string false "number of elements"
// @Success 200 {object} models.EmailsList
// @Router /email/search [get]
func (h *emailHandlers) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "emailHandlers.Search")
		defer span.Finish()
		searchRequests.Inc()

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			h.log.Errorf("strconv.Atoi: %v", err)
			errorRequests.Inc()
			return httpErrors.ErrorCtxResponse(c, httpErrors.BadRequest)
		}
		size, err := strconv.Atoi(c.QueryParam("size"))
		if err != nil {
			h.log.Errorf("strconv.Atoi: %v", err)
			errorRequests.Inc()
			return httpErrors.ErrorCtxResponse(c, httpErrors.BadRequest)
		}

		pq := utils.NewPaginationQuery(size, page)

		res, err := h.emailUC.Search(ctx, c.QueryParam("search"), pq)
		if err != nil {
			h.log.Errorf("strconv.Atoi: %v", err)
			errorRequests.Inc()
			return httpErrors.ErrorCtxResponse(c, httpErrors.BadRequest)
		}

		successRequests.Inc()
		return c.JSON(http.StatusOK, res)
	}
}
