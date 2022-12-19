package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (s *merchantService) GetMerchantProfile(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
	data, err := s.repo.GetMerchantProfile(ctx, merchantID)
	if err != nil {
		if err != customerrors.ErrRecordNotFound {
			s.log.Warningln("[GetMerchantProfile] Failed on running the repository:", err.Error())

		}
		return nil, err
	}

	if data.DisplayProfileURL == nil {
		data.DisplayProfileURL = &s.config.DefaultProfilePictureURL
	}
	return data, nil
}
