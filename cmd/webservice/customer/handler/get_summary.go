package handler

import (
	"log"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *customerHandler) GetSummary() echo.HandlerFunc {
	return func(c echo.Context) error {
		customerCtx := authutils.CustomerFromRequestContext(c)
		if customerCtx == nil {
			log.Println("[GetCustomerDetails] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		summary, err := h.service.GetSummary(c.Request().Context(), customerCtx.ID)
		if err != nil {
			if err != customerrors.ErrRecordNotFound {
				log.Println("[GetCustomerDetails] Couldn't extract user account from context")
			}
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: *summary,
		})
	}
}
