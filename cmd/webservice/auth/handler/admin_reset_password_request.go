package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *authHandler) AdminResetPasswordRequest() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.AdminResetPasswordRequest
		err := c.Bind(&req)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrBadRequest,
			})
		}

		err = h.validator.StructCtx(c.Request().Context(), req)
		if err != nil {
			errStr := h.validator.TranslateValidatorError(err)
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    customerrors.ErrBadRequest,
				Detail: errStr,
			})
		}

		err = h.service.AdminResetPasswordRequest(c.Request().Context(), req.Email)
		if err != nil {
			if err == customerrors.ErrRecordNotFound {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: err,
				})
			}

			if err == customerrors.ErrUnauthorized {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: err,
				})
			}

			h.log.Warningln("[CustomerResetPasswordRequest] Error while calling the service function:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: "Email Sent!",
		})
	}
}
