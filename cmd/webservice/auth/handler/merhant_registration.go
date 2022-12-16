package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *authHandler) RegisterMerchant() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.MerchantRegisterRequest
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

		// Get error on nested struct
		var nestedErrs []error
		for _, bank := range req.MerchantBank {
			err = h.validator.StructCtx(c.Request().Context(), bank)
			nestedErrs = append(nestedErrs, err)
		}

		if nestedErrs[0] != nil {
			var nestedDetails []map[string]interface{}
			for index, nestedErr := range nestedErrs {
				errStr := h.validator.TranslateValidatorError(nestedErr)
				nestedDetails = append(nestedDetails, map[string]interface{}{
					"bank_index":   index,
					"error_detail": errStr,
				})
			}

			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    customerrors.ErrBadRequest,
				Detail: nestedDetails,
			})
		}

		err = h.service.MerchantRegistration(c.Request().Context(), &req)
		if err != nil {
			if err != customerrors.ErrAccountDuplicated ||
				err != customerrors.ErrMerchantNameDuplicated ||
				err != customerrors.ErrUsernameDuplicated {
				h.log.Warningln("[RegisterMerchant] Error while calling the service function")
			}
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Code: 201,
			Data: "Merchant Created!",
		})
	}
}
