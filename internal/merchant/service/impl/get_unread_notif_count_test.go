package impl

import (
	"context"
	"testing"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/go-playground/assert/v2"
)

func (s merchantServiceSuite) TestGetUnreadNotifCount() {
	tests := []struct {
		name     string
		repo     funcGetUnreadNotifCountMerchant
		expected struct {
			resp int
			err  error
		}
	}{
		{
			name: "Valid",
			repo: func(ctx context.Context, MerchantID int) (int, error) {
				return 1, nil
			},
			expected: struct {
				resp int
				err  error
			}{
				resp: 1,
				err:  nil,
			},
		},
		{
			name: "Error because repo return error",
			repo: func(ctx context.Context, MerchantID int) (int, error) {
				return 0, customerrors.ErrInternalServer
			},
			expected: struct {
				resp int
				err  error
			}{
				resp: 0,
				err:  customerrors.ErrInternalServer,
			},
		},
	}

	for _, test := range tests {
		s.service.repoNotif = mockNotificationRepository{
			funcGetUnreadNotifCountMerchant: test.repo,
		}

		resp, err := s.service.GetUnreadNotifCount(context.Background(), 1)

		s.Suite.T().Run(test.name, func(t *testing.T) {
			assert.Equal(t, err, test.expected.err)
			assert.Equal(t, resp, test.expected.resp)
		})
	}
}
