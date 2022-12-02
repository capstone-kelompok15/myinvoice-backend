package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/passwordutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/randomutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/tokenutils"
)

func (s *customerService) LoginAdmin(ctx context.Context, req *dto.AdminLoginRequest) (*dto.AdminLoginResponse, error) {
	req.Password = passwordutils.HashPassword(req.Password)

	adminContext, err := s.repo.LoginAdmin(ctx, req)
	if err != nil {
		if err != customerrors.ErrRecordNotFound {
			s.log.Warningln("[CustomerLogin] Error while authorize customer login", err.Error())
		}
		return nil, customerrors.ErrUnauthorized
	}

	var loginResponse dto.AdminLoginResponse
	loginResponse.AccessToken, err = tokenutils.NewAccessToken(s.config.JWTSecretKey, adminContext)
	if err != nil {
		s.log.Warningln("[CustomerLogin] Error while creating access token", err.Error())
		return nil, customerrors.ErrUnauthorized
	}

	refreshToken := randomutils.GenerateNRandomString(64)
	loginResponse.RefreshToken = refreshToken

	return nil, nil
}
