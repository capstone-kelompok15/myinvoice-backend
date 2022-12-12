package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type NotificationRepository interface {
	GetTitleID(title string) (int, error)
	CheckNotifCustomerExist(ID int) error
	CheckNotifMerchantExist(ID int) error
	GetAllNotificationCustomer(ctx context.Context, customerID int, req *dto.NotificationRequest) (*[]dto.CustomerNotificationDB, error)
	GetAllNotificationMerchant(ctx context.Context, merchantID int, req *dto.NotificationRequest) (*[]dto.MerchantNotificationDB, error)
	CreateNotificationCustomer(ctx context.Context, req *dto.CreateNotification) error
	CreateNotificationMerchant(ctx context.Context, req *dto.CreateNotification) error
	MarkNotifCustomerAsRead(ctx context.Context, NotifID int, CustomerID int) error
	MarkNotifMerchantAsRead(ctx context.Context, NotifID int, MerchantID int) error
}
