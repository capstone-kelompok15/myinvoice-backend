package handler

import (
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *invoiceHandler) UpdatePaymentMethod() echo.HandlerFunc {
	return func(c echo.Context) error {

		req := struct {
			InvoiceID      int `param:"invoice_id" validate:"required"`
			MerchantBankID int `json:"merchant_bank_id" validate:"required"`
		}{}
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
		err = h.service.UpdatePaymentMethod(c.Request().Context(), req.InvoiceID, req.MerchantBankID)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Code: 200,
			Data: "Update payment method success",
		})
	}
}
