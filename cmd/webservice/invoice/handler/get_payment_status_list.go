package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *invoiceHandler) GetPaymentStatusList() echo.HandlerFunc {
	return func(c echo.Context) error {
		statusList, err := h.service.GetPaymentStatusList(c.Request().Context())
		if err != nil {
			h.log.Warningln("[GetPaymentStatusList] Error on running the repository:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}
		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: *statusList,
		})
	}
}
