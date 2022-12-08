package service

import "github.com/labstack/echo/v4"

type Middleware interface {
	CustomerMustAuthorized() echo.MiddlewareFunc
	MustAuthorized() echo.MiddlewareFunc
}
