package delivery

import "github.com/labstack/echo/v4"

// HTTPDelivery interface
type HTTPDelivery interface {
	Create() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Search() echo.HandlerFunc
}
