package handler

import (
	"log"
	"strconv"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *customerHandler) MarkNotifCustomerAsRead() echo.HandlerFunc {
	return func(c echo.Context) error {
		customerCtx := authutils.CustomerFromRequestContext(c)
		if customerCtx == nil {
			log.Println("[MarkNotifCustomerAsRead] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}
		NotifID, _ := strconv.Atoi(c.Param("id"))
		if NotifID == 0 {

			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrBadRequest,
			})
		}

		err := h.service.MarkNotifCustomerAsRead(c.Request().Context(), NotifID, customerCtx.ID)
		if err != nil {
			log.Println("[MarkNotifCustomerAsRead] Couldn't Mark Notif Customer As Read")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: "Mark notif customer as read success",
		})
	}
}
