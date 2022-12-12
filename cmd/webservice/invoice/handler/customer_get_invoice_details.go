package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *invoiceHandler) CustomerGetDetailInvoiceByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.GetDetailsInvoicesRequest
		customerCtx := authutils.CustomerFromRequestContext(c)

		if customerCtx == nil {
			h.log.Warningln("[CustomerGetDetailInvoiceByID] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}
		req.CustomerID = customerCtx.ID

		err := c.Bind(&req)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrBadRequest,
			})
		}

		invoice, err := h.service.GetDetailInvoiceByID(c.Request().Context(), &req)
		if err != nil {
			if err != customerrors.ErrRecordNotFound {
				h.log.Warningln("[CustomerGetDetailInvoiceByID] Error on running the service:", err.Error())
			}
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: *invoice,
		})
	}
}
