package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/passwordutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/tokenutils"
)

func (s *customerService) CustomerLogin(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerAccessToken, error) {
	req.Password = passwordutils.HashPassword(req.Password)

	customerContext, err := s.repo.AuthorizeCustomerLogin(ctx, req)
	if err != nil {
		s.log.Warningln("[CustomerLogin] Error while authorize customer login", err.Error())
		return nil, customerrors.ErrUnauthorized
	}

	err = s.repo.InvalidCustomerAccessToken(ctx, customerContext.ID)
	if err != nil {
		s.log.Warningln("[CustomerLogin] Error while calling the repo function", err.Error())
		return nil, err
	}

	deviceID := passwordutils.HashPassword(req.DeviceID)
	err = s.repo.InsertCustomerAccessToken(ctx, customerContext.ID, deviceID)
	if err != nil {
		s.log.Warningln("[CustomerLogin] Error while calling the repo function", err.Error())
		return nil, err
	}

	accessTokenStr, err := tokenutils.GenerateCustomerAccessToken(&tokenutils.CustomerAccessTokenParams{
		DeviceInformation: req.DeviceID,
		UserInformation:   customerContext,
		Config:            &s.config.CustomerToken,
	})
	if err != nil {
		s.log.Warningln("[CustomerLogin] Error while generating customer access token", err.Error())
		return nil, err
	}

	accessToken := dto.CustomerAccessToken{
		AccessToken: *accessTokenStr,
	}

	return &accessToken, nil
}
