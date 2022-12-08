package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type CustomerRepository interface {
	GetCustomerDetail(ctx context.Context, customerID int) (*dto.CustomerDetails, error)
	UpdateProfilePicture(ctx context.Context, userID *int, newProfilePictureURL *string) error
	UpdateCustomer(ctx context.Context, customerID *int, newData *dto.CustomerUpdateRequest) error
}
