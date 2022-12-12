package service

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type MerchantService interface {
	GetDashboard(ctx context.Context, merchantID int) (*dto.MerchantDashbaord, error)
	GetAllNotificationMerchant(ctx context.Context, merchantID int, req *dto.NotificationRequest) (*[]dto.NotificationRespond, error)
	MarkNotifMerchantAsRead(ctx context.Context, NotifID int, MerchantID int) error
}
