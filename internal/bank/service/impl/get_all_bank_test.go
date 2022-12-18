package impl

import (
	"context"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func (s bankServiceSuite) TestGetAllBank() {
	tests := []struct {
		name          string
		repo          funcGetAllBank
		expected      *[]dto.BankResponse
		expectedError error
	}{
		{
			name: "Valid",
			repo: func(ctx context.Context) (*[]dto.BankResponse, error) {
				return &[]dto.BankResponse{
					{
						ID:       1,
						BankName: "BCA",
						Code:     "BCA",
					},
				}, nil
			},
			expected: &[]dto.BankResponse{
				{
					ID:       1,
					BankName: "BCA",
					Code:     "BCA",
				},
			},
			expectedError: nil,
		},
		{
			name: "Repo return error",
			repo: func(ctx context.Context) (*[]dto.BankResponse, error) {
				return nil, customerrors.ErrInternalServer
			},
			expected:      nil,
			expectedError: customerrors.ErrInternalServer,
		},
		{
			name: "Repo return record not found",
			repo: func(ctx context.Context) (*[]dto.BankResponse, error) {
				return nil, customerrors.ErrRecordNotFound
			},
			expected:      nil,
			expectedError: customerrors.ErrRecordNotFound,
		},
	}

	for _, test := range tests {
		s.service.repo = mockBankRepository{
			funcGetAllBank: test.repo,
		}

		res, err := s.service.GetAllBank(context.Background())

		s.Suite.T().Run(test.name, func(t *testing.T) {
			assert.Equal(t, err, test.expectedError)
			assert.Equal(t, test.expected, res)
		})
	}
}
