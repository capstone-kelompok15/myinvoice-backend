package service

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type MerchantService interface {
	GetAllNotificationMerchant(ctx context.Context, merchantID int, req *dto.NotificationRequest) (*[]dto.NotificationRespond, error)
	MarkNotifMerchantAsRead(ctx context.Context, NotifID int, MerchantID int) error
	GetDashboard(ctx context.Context, merchantID int) (*dto.MerchantDashboard, error)
	GetMerchantBank(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error)
	UpdateMerchantBank(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error
	CreateMerchantBank(ctx context.Context, merchantID int, req *dto.MerchantBankData) error
	UpdateProfilePicture(ctx context.Context, userID *int, newProfilePictureURL *string) (*string, error)
}
