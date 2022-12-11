package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (s *merchantService) CreateMerchantBank(ctx context.Context, merchantID int, req *dto.MerchantBankData) error {
	err := s.repo.CreateMerchantBank(ctx, merchantID, req)
	if err != nil {
		s.log.Warningln("[CrateMerchantBank] Error of running the repository:", err.Error())
		return err
	}
	return nil
}
