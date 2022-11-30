package handler

import (
	"database/sql"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *customerHandler) RefreshEmailVerificationCode() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.CustomerRefreshEmailVerificationcode
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

		err = h.service.RefreshEmailVerificationCode(c.Request().Context(), req.Email)
		if err != nil {

			if err == sql.ErrNoRows {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: customerrors.ErrRecordNotFound,
				})
			}

			h.log.Warningln("[RefreshEmailVerificationCode] Error while calling the service function")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Code: 201,
			Data: "Sent new code!",
		})
	}
}
