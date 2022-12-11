package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type MerchantRepository interface {
	GetDashboardInvoiceOverview(ctx context.Context, merchantID int) (*dto.OverviewMerchantDashboard, error)
	GetDashboardRecentInvoices(ctx context.Context, merchantID int) (*[]dto.RecentInvoiceMerchantDashboard, error)
	GetDashboardRecentPayments(ctx context.Context, merchantID int) (*[]dto.RecentPaymentMerchantDashboard, error)
	GetMerchantBank(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error)
	UpdateMerchantBank(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error
	ValidateMerchantBank(ctx context.Context, merchantID int, merchantBankID int) error
	CreateMerchantBank(ctx context.Context, merchantID int, req *dto.MerchantBankData) error
}
