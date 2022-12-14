package handler

import (
	"os"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *invoiceHandler) UploadPayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		customerCtx := authutils.CustomerFromRequestContext(c)
		if customerCtx == nil {
			h.log.Warningln("[UpdateCustomer] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		var req struct {
			InvoiceID int `param:"invoice_id" validate:"required"`
		}

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

		paymentPicture, err := c.FormFile("payment")
		if paymentPicture == nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrBadRequest,
			})
		}

		if err != nil {
			h.log.Warningln("[UploadPayment] error while getting the form file:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		profilePictureFileName, err := httputils.HandleFileForm(paymentPicture)
		if err != nil {
			h.log.Warningln("[UploadPayment] error while creating the file:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}
		defer os.Remove(*profilePictureFileName)

		message, err := h.service.UploadPayment(c.Request().Context(), customerCtx.ID, req.InvoiceID, *profilePictureFileName)
		if err != nil {
			if err == customerrors.ErrRecordNotFound {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: customerrors.ErrRecordNotFound,
				})
			}
			h.log.Warningln("[UploadPayment] error on running the service fuction:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		h.websocketPool.Message <- message

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: "Upload Successful!",
		})
	}
}
