package impl

import (
	"context"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/go-playground/assert/v2"
)

func (s merchantServiceSuite) TestGetDashboard() {
	tests := []struct {
		name string
		repo struct {
			funcGetDashboardInvoiceOverview
			funcGetDashboardRecentInvoices
			funcGetDashboardRecentPayments
		}
		expected struct {
			resp *dto.MerchantDashboard
			err  error
		}
	}{
		{
			name: "Valid",
			repo: struct {
				funcGetDashboardInvoiceOverview
				funcGetDashboardRecentInvoices
				funcGetDashboardRecentPayments
			}{
				funcGetDashboardInvoiceOverview: func(ctx context.Context, merchantID int) (*dto.OverviewMerchantDashboard, error) {
					return &dto.OverviewMerchantDashboard{}, nil
				},
				funcGetDashboardRecentInvoices: func(ctx context.Context, merchantID int) (*[]dto.RecentInvoiceMerchantDashboard, error) {
					return &[]dto.RecentInvoiceMerchantDashboard{}, nil
				},
				funcGetDashboardRecentPayments: func(ctx context.Context, merchantID int) (*[]dto.RecentPaymentMerchantDashboard, error) {
					return &[]dto.RecentPaymentMerchantDashboard{}, nil
				},
			},
			expected: struct {
				resp *dto.MerchantDashboard
				err  error
			}{
				resp: &dto.MerchantDashboard{
					OverviewMerchantDashboard:      dto.OverviewMerchantDashboard{},
					RecentInvoiceMerchantDashboard: []dto.RecentInvoiceMerchantDashboard{},
					RecentPaymentMerchantDashboard: []dto.RecentPaymentMerchantDashboard{},
				},
				err: nil,
			},
		},
		{
			name: "Error because on of the repo return error",
			repo: struct {
				funcGetDashboardInvoiceOverview
				funcGetDashboardRecentInvoices
				funcGetDashboardRecentPayments
			}{
				funcGetDashboardInvoiceOverview: func(ctx context.Context, merchantID int) (*dto.OverviewMerchantDashboard, error) {
					return nil, customerrors.ErrInternalServer
				},
				funcGetDashboardRecentInvoices: func(ctx context.Context, merchantID int) (*[]dto.RecentInvoiceMerchantDashboard, error) {
					return &[]dto.RecentInvoiceMerchantDashboard{}, nil
				},
				funcGetDashboardRecentPayments: func(ctx context.Context, merchantID int) (*[]dto.RecentPaymentMerchantDashboard, error) {
					return &[]dto.RecentPaymentMerchantDashboard{}, nil
				},
			},
			expected: struct {
				resp *dto.MerchantDashboard
				err  error
			}{
				resp: nil,
				err:  customerrors.ErrInternalServer,
			},
		},
	}

	for _, test := range tests {
		s.service.repo = mockMerchantRepository{
			funcGetDashboardInvoiceOverview: test.repo.funcGetDashboardInvoiceOverview,
			funcGetDashboardRecentInvoices:  test.repo.funcGetDashboardRecentInvoices,
			funcGetDashboardRecentPayments:  test.repo.funcGetDashboardRecentPayments,
		}

		resp, err := s.service.GetDashboard(context.Background(), 1)

		s.Suite.T().Run(test.name, func(t *testing.T) {
			assert.Equal(t, err, test.expected.err)
			if err != nil {
				assert.Equal(t, resp, test.expected.resp)
			} else {
				assert.Equal(t, resp.OverviewMerchantDashboard, test.expected.resp.OverviewMerchantDashboard)
				assert.Equal(t, resp.RecentInvoiceMerchantDashboard, test.expected.resp.RecentInvoiceMerchantDashboard)
				assert.Equal(t, resp.RecentPaymentMerchantDashboard, test.expected.resp.RecentPaymentMerchantDashboard)
			}
		})
	}
}
