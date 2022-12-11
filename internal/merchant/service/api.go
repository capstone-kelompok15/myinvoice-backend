package service

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type MerchantService interface {
	GetDashboard(ctx context.Context, merchantID int) (*dto.MerchantDashboard, error)
	GetMerchantBank(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error)
	UpdateMerchantBank(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error
}
