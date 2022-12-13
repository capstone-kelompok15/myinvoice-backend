package handler

import (
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *merchantHandler) GetMerchantProfile() echo.HandlerFunc {
	return func(c echo.Context) error {

		adminCtx := authutils.AdminContextFromRequestContext(c)
		if adminCtx == nil {
			h.log.Warningln("[UpdateMerchantBank] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		data, err := h.service.GetMerchantProfile(c.Request().Context(), adminCtx.ID)
		if err != nil {
			h.log.Warningln("[GetMerchantProfile] Error while calling the service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: *data,
		})
	}
}
