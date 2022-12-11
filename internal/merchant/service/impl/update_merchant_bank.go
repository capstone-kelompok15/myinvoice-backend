package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (s *merchantService) UpdateMerchantBank(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error {
	err := s.repo.ValidateMerchantBank(ctx, req.MerchantID, req.MerchantBankID)
	if err != nil {
		if err == customerrors.ErrRecordNotFound {
			return customerrors.ErrRecordNotFound
		}
		s.log.Warningln("[UpdateMerchantBank] Error while validate merchant bank")
		return err
	}

	err = s.repo.UpdateMerchantBank(ctx, req)
	if err != nil {
		if err == customerrors.ErrRecordNotFound {
			return err
		}
		s.log.Warningln("[UpdateMerchantBank] Error while executing update bank merchant")
		return err
	}

	return nil
}
