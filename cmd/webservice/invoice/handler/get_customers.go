package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *invoiceHandler) GetCustomers() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := dto.GetMerchantCustomerList{
			PaginationFilter: nil,
			MerchantID:       0,
		}

		adminCtx := authutils.AdminContextFromRequestContext(c)
		if adminCtx == nil {
			h.log.Warningln("[GetCustomers] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}
		req.MerchantID = adminCtx.MerchantID

		var err error
		req.PaginationFilter = &dto.PaginationFilter{}
		req.PaginationFilter.Offset, req.PaginationFilter.Limit, err = httputils.GetPaginationMandatoryParams(c, true)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrBadRequest,
				Detail: []string{
					"limit and offset couldn't be empty",
				},
			})
		}

		customers, count, err := h.service.GetCustomers(c.Request().Context(), &req)
		if err != nil {
			h.log.Warningln("[GetCustomers] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: map[string]interface{}{
				"customers": *customers,
				"pagination": map[string]int{
					"qty": count,
				},
			},
		})
	}
}
