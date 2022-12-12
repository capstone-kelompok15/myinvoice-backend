package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/dateutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *invoiceHandler) CreateInvoice() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.CreateInvoiceRequest
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
			h.log.Warningln("[CreateMerchantBank] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		var nestedErrs []error
		for _, invoiceDetail := range req.InvoiceDetails {
			err = h.validator.StructCtx(c.Request().Context(), invoiceDetail)
			nestedErrs = append(nestedErrs, err)
		}

		if nestedErrs[0] != nil {
			var nestedDetails []map[string]interface{}
			for index, nestedErr := range nestedErrs {
				errStr := h.validator.TranslateValidatorError(nestedErr)
				nestedDetails = append(nestedDetails, map[string]interface{}{
					"invoice_index": index,
					"error_detail":  errStr,
				})
			}

			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    customerrors.ErrBadRequest,
				Detail: nestedDetails,
			})
		}

		dueAt, err := dateutils.StringToDate(req.DueAtRequest)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}
		req.DueAt = *dueAt

		err = h.service.CreateInvoice(c.Request().Context(), adminCtx.MerchantID, &req)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: "New invoice created!",
		})
	}
}
