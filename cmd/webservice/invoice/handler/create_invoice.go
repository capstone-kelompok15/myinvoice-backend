package handler

import (
	"time"

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

		dueAt, err := dateutils.StringToDate(req.DueAtRequest)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}
		req.DueAt = *dueAt

		if dueAt.Before(time.Now().AddDate(0, 0, -1)) {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrBadRequest,
				Detail: []string{
					"Due At value can be smaller than today",
				},
			})
		}

		var isErr bool
		var nestedErrs []error
		for _, invoiceDetail := range req.Items {
			err = h.validator.StructCtx(c.Request().Context(), invoiceDetail)
			if err != nil {
				isErr = true
			}
			nestedErrs = append(nestedErrs, err)
		}

		if isErr {
			var nestedDetails []map[string]interface{}
			for index, nestedErr := range nestedErrs {
				if nestedErr != nil {
					errStr := h.validator.TranslateValidatorError(nestedErr)
					nestedDetails = append(nestedDetails, map[string]interface{}{
						"invoice_index": index,
						"error_detail":  errStr,
					})
				}
			}

			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    customerrors.ErrBadRequest,
				Detail: nestedDetails,
			})
		}

		adminCtx := authutils.AdminContextFromRequestContext(c)
		if adminCtx == nil {
			h.log.Warningln("[CreateInvoice] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		err = h.service.CreateInvoice(c.Request().Context(), adminCtx.MerchantID, &req)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Code: 201,
			Data: "New invoice created!",
		})
	}
}
