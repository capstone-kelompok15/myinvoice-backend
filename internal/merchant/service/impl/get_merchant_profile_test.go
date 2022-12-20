package impl

import (
	"context"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/go-playground/assert/v2"
)

func (s merchantServiceSuite) TestGetMerchantProfile() {
	tests := []struct {
		name     string
		repo     funcGetMerchantProfile
		expected struct {
			resp *dto.MerchantProfileResponse
			err  error
		}
	}{
		{
			name: "Valid",
			repo: func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
				return &dto.MerchantProfileResponse{}, nil
			},
			expected: struct {
				resp *dto.MerchantProfileResponse
				err  error
			}{
				resp: &dto.MerchantProfileResponse{
					DisplayProfileURL: &s.service.config.DefaultProfilePictureURL,
				},
				err: nil,
			},
		},
		{
			name: "Error because repo return error",
			repo: func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
				return nil, customerrors.ErrInternalServer
			},
			expected: struct {
				resp *dto.MerchantProfileResponse
				err  error
			}{
				resp: nil,
				err:  customerrors.ErrInternalServer,
			},
		},
		{
			name: "Error because record not found",
			repo: func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
				return nil, customerrors.ErrRecordNotFound
			},
			expected: struct {
				resp *dto.MerchantProfileResponse
				err  error
			}{
				resp: nil,
				err:  customerrors.ErrRecordNotFound,
			},
		},
	}

	for _, test := range tests {
		s.service.repo = mockMerchantRepository{
			funcGetMerchantProfile: test.repo,
		}

		resp, err := s.service.GetMerchantProfile(context.Background(), 1)

		s.Suite.T().Run(test.name, func(t *testing.T) {
			assert.Equal(t, err, test.expected.err)
			assert.Equal(t, resp, test.expected.resp)
		})
	}
}
