package impl

import (
	"context"
	"testing"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/go-playground/assert/v2"
)

func (s merchantServiceSuite) TestMarkNotifMerchantAsRead() {
	tests := []struct {
		name string
		repo struct {
			funcCheckNotifMerchantExist
			funcMarkNotifMerchantAsRead
		}
		expected struct {
			error
		}
	}{
		{
			name: "Valid",
			repo: struct {
				funcCheckNotifMerchantExist
				funcMarkNotifMerchantAsRead
			}{
				funcCheckNotifMerchantExist: func(ID int) error {
					return nil
				},
				funcMarkNotifMerchantAsRead: func(ctx context.Context, NotifID, MerchantID int) error {
					return nil
				},
			},
			expected: struct{ error }{
				error: nil,
			},
		},
		{
			name: "Error because repo return error",
			repo: struct {
				funcCheckNotifMerchantExist
				funcMarkNotifMerchantAsRead
			}{
				funcCheckNotifMerchantExist: func(ID int) error {
					return customerrors.ErrInternalServer
				},
				funcMarkNotifMerchantAsRead: func(ctx context.Context, NotifID, MerchantID int) error {
					return nil
				},
			},
			expected: struct{ error }{
				error: customerrors.ErrInternalServer,
			},
		},
	}

	for _, test := range tests {
		s.service.repoNotif = mockNotificationRepository{
			funcCheckNotifMerchantExist: test.repo.funcCheckNotifMerchantExist,
			funcMarkNotifMerchantAsRead: test.repo.funcMarkNotifMerchantAsRead,
		}

		err := s.service.MarkNotifMerchantAsRead(context.Background(), 1, 1)

		s.Suite.T().Run(test.name, func(t *testing.T) {
			assert.Equal(t, err, test.expected.error)
		})
	}
}
