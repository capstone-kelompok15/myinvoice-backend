package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *merchantHandler) CreateMerchantBank() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.MerchantBankData
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

		adminCtx := authutils.AdminContextFromRequestContext(c)
		if adminCtx == nil {
			h.log.Warningln("[CreateMerchantBank] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		err = h.service.CreateMerchantBank(c.Request().Context(), adminCtx.MerchantID, &req)
		if err != nil {
			h.log.Warningln("[CreateMerchantBank] Error on running the service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Code: 201,
			Data: "New Merchant Bank Created!",
		})
	}
}
