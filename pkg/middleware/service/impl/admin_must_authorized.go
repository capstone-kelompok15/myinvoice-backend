package impl

import (
	"strings"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/tokenutils"
	"github.com/labstack/echo/v4"
)

func (m *middleware) AdminMustAuthorized() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err:    customerrors.ErrBadRequest,
					Detail: []string{"Authorization header value couldn't be empty"},
				})
			}

			splitted := strings.SplitAfter(authorization, "Bearer ")
			if len(splitted) != 2 {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err:    customerrors.ErrBadRequest,
					Detail: []string{"Bearer format is not valid"},
				})
			}

			if splitted[1] == "" {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err:    customerrors.ErrBadRequest,
					Detail: []string{"Bearer value is couldn't empty"},
				})
			}

			accessToken := splitted[1]

			user, err := tokenutils.ValidateAccessToken(&m.config.JWTConfig, accessToken)
			if err != nil {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: customerrors.ErrUnauthorized,
				})
			}

			c.Set(dto.AccountCTXKey, user)

			return next(c)
		}
	}
}
