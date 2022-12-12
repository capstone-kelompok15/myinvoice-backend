package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (s *merchantService) GetDashboard(ctx context.Context, merchantID int) (*dto.MerchantDashboard, error) {
	errChan := make(chan error, 3)
	var merchantDashboard dto.MerchantDashboard

	go func() {
		overviewMerchantDashboard, err := s.repo.GetDashboardInvoiceOverview(ctx, merchantID)
		if overviewMerchantDashboard != nil {
			merchantDashboard.OverviewMerchantDashboard = *overviewMerchantDashboard
		}
		errChan <- err
	}()

	go func() {
		recentInvoiceMerchantDashboard, err := s.repo.GetDashboardRecentInvoices(ctx, merchantID)
		merchantDashboard.RecentInvoiceMerchantDashboard = *recentInvoiceMerchantDashboard
		errChan <- err
	}()

	go func() {
		recentPaymentMerchantDashboard, err := s.repo.GetDashboardRecentPayments(ctx, merchantID)
		merchantDashboard.RecentPaymentMerchantDashboard = *recentPaymentMerchantDashboard
		errChan <- err
	}()

	for i := 0; i < 3; i++ {
		err := <-errChan
		if err != nil {
			s.log.Warningln("[GetDashboard] Failed on running the repository:", err.Error())
			return nil, err
		}
		continue
	}

	return &merchantDashboard, nil
}
