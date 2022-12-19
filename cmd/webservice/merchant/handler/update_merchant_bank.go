package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *merchantHandler) UpdateMerchantBank() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.UpdateMerchantBankDataRequest
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

		adminCtx := authutils.AdminContextFromRequestContext(c)
		if adminCtx == nil {
			h.log.Warningln("[UpdateMerchantBank] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}
		req.MerchantID = adminCtx.MerchantID

		err = h.service.UpdateMerchantBank(c.Request().Context(), &req)
		if err != nil {
			if err == customerrors.ErrRecordNotFound {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: customerrors.ErrRecordNotFound,
				})
			}
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: "Update Merchant Bank Success!",
		})
	}
}
