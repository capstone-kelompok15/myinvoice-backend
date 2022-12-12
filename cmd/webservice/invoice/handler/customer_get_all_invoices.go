package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *invoiceHandler) CustomerGetAllInvoices() echo.HandlerFunc {
	return func(c echo.Context) error {
		filter := dto.GetAllInvoicesParam{
			DateFilter:       nil,
			PaginationFilter: nil, // required
			CustomerID:       0,
			PaymentStatusID:  0,
		}

		customerCtx := authutils.CustomerFromRequestContext(c)
		if customerCtx == nil {
			h.log.Warningln("[GetAllNotificationCustomer] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}
		filter.CustomerID = customerCtx.ID

		err := c.Bind(&filter)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrBadRequest,
			})
		}

		filter.PaginationFilter = &dto.PaginationFilter{}
		filter.PaginationFilter.Offset, filter.PaginationFilter.Limit, err = httputils.GetPaginationMandatoryParams(c, true)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrBadRequest,
				Detail: []string{
					"limit and offset couldn't be empty",
				},
			})
		}

		if c.QueryParam("start_date") != "" {
			filter.DateFilter = &dto.DateFilter{}
			filter.DateFilter.StartDate, filter.DateFilter.EndDate = httputils.GetDateQueryMandatoryParams(c)
		} else if c.QueryParam("end_date") != "" {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrBadRequest,
				Detail: []string{
					"start date couldn't be empty if end date filled",
				},
			})
		}

		invoices, count, err := h.service.GetAllInvoice(c.Request().Context(), &filter)
		if err != nil {
			h.log.Warningln("[MerchantGetAllInvoices] Service error:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: map[string]interface{}{
				"pagination": map[string]int{
					"qty": count,
				},
				"invoices": invoices,
			},
		})
	}
}
