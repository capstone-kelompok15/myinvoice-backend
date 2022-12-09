package service

import "github.com/labstack/echo/v4"

type Middleware interface {
	CustomerMustAuthorized() echo.MiddlewareFunc
	AdminMustAuthorized() echo.MiddlewareFunc
}
