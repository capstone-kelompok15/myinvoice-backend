package impl

import (
	"context"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/go-playground/assert/v2"
)

func (s merchantServiceSuite) TestCreateMerchantBank() {
	tests := []struct {
		name     string
		repo     funcCreateMerchantBank
		expected struct {
			error
		}
	}{
		{
			name: "Valid",
			repo: func(ctx context.Context, merchantID int, req *dto.MerchantBankData) error {
				return nil
			},
			expected: struct{ error }{
				error: nil,
			},
		},
		{
			name: "Error because repo returning error",
			repo: func(ctx context.Context, merchantID int, req *dto.MerchantBankData) error {
				return customerrors.ErrInternalServer
			},
			expected: struct{ error }{
				error: customerrors.ErrInternalServer,
			},
		},
	}

	for _, test := range tests {
		s.service.repo = mockMerchantRepository{
			funcCreateMerchantBank: test.repo,
		}

		err := s.service.CreateMerchantBank(context.Background(), 1, nil)

		s.Suite.T().Run(test.name, func(t *testing.T) {
			assert.Equal(t, err, test.expected.error)
		})
	}
}
