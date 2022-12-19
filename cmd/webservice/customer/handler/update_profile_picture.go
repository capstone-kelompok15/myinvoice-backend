package handler

import (
	"os"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func (h *customerHandler) UpdateCustomerProfilePicture() echo.HandlerFunc {
	return func(c echo.Context) error {
		customerCtx := authutils.CustomerFromRequestContext(c)
		if customerCtx == nil {
			h.log.Warningln("[GetCustomerDetails] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		profilePicture, err := c.FormFile("profile_picture")
		if err != nil || profilePicture == nil {
			h.log.Warningln("[GetCustomerDetails] error while getting the form file:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrBadRequest,
			})
		}

		profilePictureFileName, err := httputils.HandleFileForm(profilePicture)
		if err != nil {
			h.log.Warningln("[GetCustomerDetails] error while creating the file:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}
		defer os.Remove(*profilePictureFileName)

		imageURL, err := h.service.UpdateProfilePicture(c.Request().Context(), &customerCtx.ID, profilePictureFileName)
		if err != nil {
			if err == customerrors.ErrRecordNotFound {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: customerrors.ErrRecordNotFound,
				})
			}
			h.log.Warningln("[GetCustomerDetails] error on running the service fuction:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: dto.UpdateProfilePictureResponse{
				ImageURL: *imageURL,
			},
		})
	}
}
