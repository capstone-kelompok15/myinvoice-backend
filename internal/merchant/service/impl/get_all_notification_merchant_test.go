package impl

import (
	"context"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/go-playground/assert/v2"
)

func (s merchantServiceSuite) TestGetAllNotificationMerchant() {
	tests := []struct {
		name     string
		expected struct {
			resp *[]dto.NotificationRespond
			err  error
		}
		repo funcGetAllNotificationMerchant
	}{
		{
			name: "Valid",
			expected: struct {
				resp *[]dto.NotificationRespond
				err  error
			}{
				resp: &[]dto.NotificationRespond{},
				err:  nil,
			},
			repo: func(ctx context.Context, merchantID int, req *dto.NotificationRequest) (*[]dto.MerchantNotificationDB, error) {
				return &[]dto.MerchantNotificationDB{}, nil
			},
		},
		{
			name: "Error because repo return error",
			expected: struct {
				resp *[]dto.NotificationRespond
				err  error
			}{
				resp: nil,
				err:  customerrors.ErrInternalServer,
			},
			repo: func(ctx context.Context, merchantID int, req *dto.NotificationRequest) (*[]dto.MerchantNotificationDB, error) {
				return nil, customerrors.ErrInternalServer
			},
		},
	}

	for _, test := range tests {
		s.service.repoNotif = mockNotificationRepository{
			funcGetAllNotificationMerchant: test.repo,
		}

		resp, err := s.service.GetAllNotificationMerchant(context.Background(), 1, &dto.NotificationRequest{})

		s.Suite.T().Run(test.name, func(t *testing.T) {
			assert.Equal(t, err, test.expected.err)
			assert.Equal(t, resp, test.expected.resp)
		})
	}
}
