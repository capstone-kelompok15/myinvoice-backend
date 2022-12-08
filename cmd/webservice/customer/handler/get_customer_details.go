package handler

import (
	"log"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *customerHandler) GetCustomerDetails() echo.HandlerFunc {
	return func(c echo.Context) error {
		customerCtx := authutils.CustomerFromRequestContext(c)
		if customerCtx == nil {
			log.Println("[GetCustomerDetails] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		customerDetails, err := h.service.GetCustomerDetails(c.Request().Context(), customerCtx)
		if err != nil {
			log.Println("[GetCustomerDetails] Couldn't get customer details")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: customerDetails,
		})
	}
}
