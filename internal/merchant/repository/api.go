package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type MerchantRepository interface {
	GetDashboardInvoiceOverview(ctx context.Context, merchantID int) (*dto.OverviewMerchantDashboard, error)
	GetDashboardRecentInvoices(ctx context.Context, merchantID int) (*[]dto.RecentInvoiceMerchantDashboard, error)
	GetDashboardRecentPayments(ctx context.Context, merchantID int) (*[]dto.RecentPaymentMerchantDashboard, error)
}
