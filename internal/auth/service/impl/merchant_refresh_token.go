package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/tokenutils"
)

func (s *authService) GenerateNewAccessToken(ctx context.Context, refreshTokens string) (*string, error) {
	refreshTokenRes, err := s.repo.GetRefreshToken(ctx, refreshTokens)
	if err != nil {
		if err == customerrors.ErrUnauthorized {
			return nil, err
		}
		s.log.Warningln("[GenerateNewAccessToken] Error while getting the refresh token:", err.Error())
		return nil, err
	}

	adminContext, err := s.repo.GetAdminContextByID(ctx, refreshTokenRes.AdminID)
	if err != nil {
		s.log.Warningln("[GenerateNewAccessToken] Error while getting the admin context:", err.Error())
		return nil, err
	}

	err = s.repo.InvalidateRefreshToken(ctx, refreshTokenRes)
	if err != nil {
		if err == customerrors.ErrUnauthorized {
			return nil, err
		}
		s.log.Warningln("[GenerateNewAccessToken] Error while getting the refresh token:", err.Error())
		return nil, err
	}

	accesToken, err := tokenutils.NewAccessToken(s.config.JWTSecretKey, &dto.AdminContext{
		ID:           refreshTokenRes.AdminID,
		MerchantID:   adminContext.MerchantID,
		MerchantName: adminContext.MerchantName,
	})
	if err != nil {
		s.log.Warningln("[GenerateNewAccessToken] Error while create new access token:", err.Error())
		return nil, err
	}

	return &accesToken, nil
}
