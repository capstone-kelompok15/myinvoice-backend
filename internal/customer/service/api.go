package service

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type CustomerService interface {
	GetCustomerDetails(ctx context.Context, req *dto.CustomerContext) (*dto.CustomerDetails, error)
	UpdateProfilePicture(ctx context.Context, userID *int, newProfilePictureURL *string) (*string, error)
	UpdateCustomer(ctx context.Context, customerID *int, newData *dto.CustomerUpdateRequest) error
}
