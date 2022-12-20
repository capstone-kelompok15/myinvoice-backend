package impl

import (
	"context"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/go-playground/assert/v2"
)

func (s merchantServiceSuite) TestUpdateMerchantProfile() {
	tests := []struct {
		name string
		repo struct {
			funcGetMerchantProfile
			funcUpdateMerchantProfile
		}
		expected struct {
			error
		}
	}{
		{
			name: "Valid",
			repo: struct {
				funcGetMerchantProfile
				funcUpdateMerchantProfile
			}{
				funcGetMerchantProfile: func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
					return nil, nil
				},
				funcUpdateMerchantProfile: func(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
					return nil
				},
			},
			expected: struct{ error }{
				error: nil,
			},
		},
		{
			name: "Error because get merchant profile return error",
			repo: struct {
				funcGetMerchantProfile
				funcUpdateMerchantProfile
			}{
				funcGetMerchantProfile: func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
					return nil, customerrors.ErrInternalServer
				},
				funcUpdateMerchantProfile: func(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
					return nil
				},
			},
			expected: struct{ error }{
				error: customerrors.ErrInternalServer,
			},
		},
		{
			name: "Error because update merchant profile return error",
			repo: struct {
				funcGetMerchantProfile
				funcUpdateMerchantProfile
			}{
				funcGetMerchantProfile: func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
					return nil, nil
				},
				funcUpdateMerchantProfile: func(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
					return customerrors.ErrInternalServer
				},
			},
			expected: struct{ error }{
				error: customerrors.ErrInternalServer,
			},
		},
		{
			name: "Error because get merchant profile return record not found",
			repo: struct {
				funcGetMerchantProfile
				funcUpdateMerchantProfile
			}{
				funcGetMerchantProfile: func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
					return nil, customerrors.ErrRecordNotFound
				},
				funcUpdateMerchantProfile: func(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
					return nil
				},
			},
			expected: struct{ error }{
				error: customerrors.ErrRecordNotFound,
			},
		},
		{
			name: "Error because update merchant profile return record not found",
			repo: struct {
				funcGetMerchantProfile
				funcUpdateMerchantProfile
			}{
				funcGetMerchantProfile: func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
					return nil, nil
				},
				funcUpdateMerchantProfile: func(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
					return customerrors.ErrRecordNotFound
				},
			},
			expected: struct{ error }{
				error: customerrors.ErrRecordNotFound,
			},
		},
	}

	for _, test := range tests {
		s.service.repo = mockMerchantRepository{
			funcGetMerchantProfile:    test.repo.funcGetMerchantProfile,
			funcUpdateMerchantProfile: test.repo.funcUpdateMerchantProfile,
		}

		id := 1
		err := s.service.UpdateMerchantProfile(context.Background(), &id, nil)

		s.Suite.T().Run(test.name, func(t *testing.T) {
			assert.Equal(t, err, test.expected.error)
		})
	}
}
