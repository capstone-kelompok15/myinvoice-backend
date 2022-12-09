package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type CustomerRepository interface {
	GetCustomerDetail(ctx context.Context, customerID int) (*dto.CustomerDetails, error)
	GetAllCustomer(ctx context.Context, req *dto.GetAllCustomerRequest) (*[]dto.GetAllCustomerRespond, error)
	UpdateProfilePicture(ctx context.Context, userID *int, newProfilePictureURL *string) error
}
