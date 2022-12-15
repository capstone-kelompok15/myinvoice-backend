package service

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type CustomerService interface {
	GetCustomerDetails(ctx context.Context, req *dto.CustomerContext) (*dto.CustomerDetails, error)
	GetAllCustomer(ctx context.Context, req *dto.GetAllCustomerRequest) (*[]dto.GetAllCustomerRespond, error)
	UpdateProfilePicture(ctx context.Context, userID *int, newProfilePictureURL *string) (*string, error)
	UpdateCustomer(ctx context.Context, customerID *int, newData *dto.CustomerUpdateRequest) error
	GetAllNotificationCustomer(ctx context.Context, customerID int, req *dto.NotificationRequest) (*[]dto.NotificationRespond, error)
	MarkNotifCustomerAsRead(ctx context.Context, NotifID int, CustomerID int) error
	GetSummary(ctx context.Context, customerID int) (*dto.CustomerSummary, error)
	GetUnreadNotifCount(ctx context.Context, CustomerID int) (int, error)
}
