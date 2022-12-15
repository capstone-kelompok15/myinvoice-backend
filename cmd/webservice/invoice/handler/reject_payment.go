package handler

import (
	"strconv"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *invoiceHandler) RejectPayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		invoiceID, err := strconv.Atoi(c.Param("invoice_id"))
		if err != nil || invoiceID == 0 {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrBadRequest,
			})
		}

		req := struct {
			Message string `json:"message" validate:"required"`
		}{}

		err = c.Bind(&req)
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

		message, err := h.service.RejectPayment(c.Request().Context(), invoiceID, req.Message)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		h.websocketPool.Message <- message

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Code: 200,
			Data: "Reject payment success",
		})
	}
}
