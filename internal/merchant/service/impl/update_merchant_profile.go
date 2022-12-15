package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (s *merchantService) UpdateMerchantProfile(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
	_, err := s.repo.GetMerchantProfile(ctx, *merchantID)
	if err != nil {
		if err == customerrors.ErrRecordNotFound {
			return customerrors.ErrRecordNotFound
		}
		s.log.Warningln("[UpdateMerchantProfile] Error while exec repo")
		return err
	}

	err = s.repo.UpdateMerchantProfile(ctx, merchantID, req)
	if err != nil {
		if err == customerrors.ErrRecordNotFound {
			return err
		}
		s.log.Warningln("[UpdateMerchantProfile] Error while executing update merchant profile repository")
		return err
	}

	return nil
}
