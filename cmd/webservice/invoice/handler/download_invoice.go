package handler

import (
	"fmt"
	"os"
	"time"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *invoiceHandler) DownloadInvoice() echo.HandlerFunc {
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

		err = h.validator.StructCtx(c.Request().Context(), req)
		if err != nil {
			errStr := h.validator.TranslateValidatorError(err)
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    customerrors.ErrBadRequest,
				Detail: errStr,
			})
		}

		filename, err := h.service.GeneratePDF(c.Request().Context(), &req, h.config.HTMlTemplatePath)
		if err != nil {
			if err != customerrors.ErrRecordNotFound {
				h.log.Warningln("[DownloadInvoice] error while calling the service:", err.Error())
			}
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}
		defer os.Remove(*filename)

		date := time.Now().String()

		return httputils.ServeFile(c, dto.ServeFileResponseParam{
			FileLocation:   *filename,
			AttachmentName: fmt.Sprintf("Laporan Pembayaran %s.pdf", date),
		})
	}
}
