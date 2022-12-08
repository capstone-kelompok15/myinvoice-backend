package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *authHandler) RegisterCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.CustomerRequest
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

		err = h.service.CustomerRegistration(c.Request().Context(), &req)
		if err != nil {
			if err == customerrors.ErrAccountDuplicated {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: customerrors.ErrAccountDuplicated,
				})
			}
			h.log.Warningln("[RegisterCustomer] Error while calling the service function")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Code: 201,
			Data: "User Created!",
		})
	}
}
