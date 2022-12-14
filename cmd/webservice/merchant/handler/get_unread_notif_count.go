package handler

import (
	"log"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *merchantHandler) GetUnreadNotifCount() echo.HandlerFunc {
	return func(c echo.Context) error {
		adminCtx := authutils.AdminContextFromRequestContext(c)
		if adminCtx == nil {
			h.log.Warningln("[GetUnreadNotifCount] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		count, err := h.service.GetUnreadNotifCount(c.Request().Context(), adminCtx.ID)
		if err != nil {
			log.Println("[GetUnreadNotifCount] Couldn't get unread notification count customer")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: struct {
				UnreadCount int `json:"unread_count"`
			}{UnreadCount: count},
		})
	}
}
