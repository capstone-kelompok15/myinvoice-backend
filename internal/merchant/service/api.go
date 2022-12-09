package service

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type MerchantService interface {
	GetDashboard(ctx context.Context, merchantID int) (*dto.MerchantDashbaord, error)
}
