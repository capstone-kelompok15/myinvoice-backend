package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (s *merchantService) GetMerchantBank(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error) {
	banks, err := s.repo.GetMerchantBank(ctx, req)
	if err != nil {
		s.log.Warningln("[GetMerchantBank] Failed on running the repository:", err.Error())
		return nil, err
	}
	return banks, nil
}
