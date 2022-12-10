package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *merchantHandler) GetMerchantBank() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.GetMerchantBankRequest
		err := c.Bind(&req)
		if err != nil {
			if err != nil {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: customerrors.ErrBadRequest,
				})
			}
		}

		err = h.validator.StructCtx(c.Request().Context(), req)
		if err != nil {
			errStr := h.validator.TranslateValidatorError(err)
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    customerrors.ErrBadRequest,
				Detail: errStr,
			})
		}

		merchantBank, err := h.service.GetMerchantBank(c.Request().Context(), &req)
		if err != nil {
			h.log.Warningln("[GetMerchantBank] Error while calling the service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: *merchantBank,
		})
	}
}
