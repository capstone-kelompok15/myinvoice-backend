package handler

import (
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *bankHandler) GetAllBank() echo.HandlerFunc {
	return func(c echo.Context) error {
		banks, err := h.service.GetAllBank(c.Request().Context())
		if err != nil {
			h.log.Warningln("[GetAllBank] Error while calling the service function")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: banks,
		})
	}
}
