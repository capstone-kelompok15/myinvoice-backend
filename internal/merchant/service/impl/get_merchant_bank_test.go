package impl

import (
	"context"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/go-playground/assert/v2"
)

func (s merchantServiceSuite) TestGetMerchantBank() {
	tests := []struct {
		name     string
		repo     funcGetMerchantBank
		expected struct {
			resp *[]dto.GetMerchantBankResponse
			err  error
		}
	}{
		{
			name: "Valid",
			repo: func(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error) {
				return &[]dto.GetMerchantBankResponse{}, nil
			},
			expected: struct {
				resp *[]dto.GetMerchantBankResponse
				err  error
			}{
				resp: &[]dto.GetMerchantBankResponse{},
				err:  nil,
			},
		},
		{
			name: "Error because service return error",
			repo: func(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error) {
				return nil, customerrors.ErrInternalServer
			},
			expected: struct {
				resp *[]dto.GetMerchantBankResponse
				err  error
			}{
				resp: nil,
				err:  customerrors.ErrInternalServer,
			},
		},
	}

	for _, test := range tests {
		s.service.repo = mockMerchantRepository{
			funcGetMerchantBank: test.repo,
		}

		resp, err := s.service.GetMerchantBank(context.Background(), nil)

		s.Suite.T().Run(test.name, func(t *testing.T) {
			assert.Equal(t, err, test.expected.err)
			assert.Equal(t, resp, test.expected.resp)
		})
	}
}
