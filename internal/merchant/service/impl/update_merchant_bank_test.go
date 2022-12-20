package impl

import (
	"context"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/go-playground/assert/v2"
)

func (s merchantServiceSuite) TestUpdateMerchantBank() {
	tests := []struct {
		name string
		repo struct {
			funcValidateMerchantBank
			funcUpdateMerchantBank
		}
		expected struct {
			error
		}
	}{
		{
			name: "Valid",
			repo: struct {
				funcValidateMerchantBank
				funcUpdateMerchantBank
			}{
				funcValidateMerchantBank: func(ctx context.Context, merchantID, merchantBankID int) error {
					return nil
				},
				funcUpdateMerchantBank: func(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error {
					return nil
				},
			},
			expected: struct{ error }{
				error: nil,
			},
		},
		{
			name: "Error because validate merchant bank return error",
			repo: struct {
				funcValidateMerchantBank
				funcUpdateMerchantBank
			}{
				funcValidateMerchantBank: func(ctx context.Context, merchantID, merchantBankID int) error {
					return customerrors.ErrInternalServer
				},
				funcUpdateMerchantBank: func(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error {
					return nil
				},
			},
			expected: struct{ error }{
				error: customerrors.ErrInternalServer,
			},
		},
		{
			name: "Error because update merchant bank return error",
			repo: struct {
				funcValidateMerchantBank
				funcUpdateMerchantBank
			}{
				funcValidateMerchantBank: func(ctx context.Context, merchantID, merchantBankID int) error {
					return nil
				},
				funcUpdateMerchantBank: func(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error {
					return customerrors.ErrInternalServer
				},
			},
			expected: struct{ error }{
				error: customerrors.ErrInternalServer,
			},
		},
		{
			name: "Error because update merchant bank return record not found",
			repo: struct {
				funcValidateMerchantBank
				funcUpdateMerchantBank
			}{
				funcValidateMerchantBank: func(ctx context.Context, merchantID, merchantBankID int) error {
					return nil
				},
				funcUpdateMerchantBank: func(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error {
					return customerrors.ErrRecordNotFound
				},
			},
			expected: struct{ error }{
				error: customerrors.ErrRecordNotFound,
			},
		},
		{
			name: "Error because validate merchant bank return record not found",
			repo: struct {
				funcValidateMerchantBank
				funcUpdateMerchantBank
			}{
				funcValidateMerchantBank: func(ctx context.Context, merchantID, merchantBankID int) error {
					return customerrors.ErrRecordNotFound
				},
				funcUpdateMerchantBank: func(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error {
					return nil
				},
			},
			expected: struct{ error }{
				error: customerrors.ErrRecordNotFound,
			},
		},
	}

	for _, test := range tests {
		s.service.repo = mockMerchantRepository{
			funcValidateMerchantBank: test.repo.funcValidateMerchantBank,
			funcUpdateMerchantBank:   test.repo.funcUpdateMerchantBank,
		}

		err := s.service.UpdateMerchantBank(context.Background(), &dto.UpdateMerchantBankDataRequest{
			MerchantID:     1,
			MerchantBankID: 1,
		})

		s.Suite.T().Run(test.name, func(t *testing.T) {
			assert.Equal(t, err, test.expected.error)
		})
	}
}
