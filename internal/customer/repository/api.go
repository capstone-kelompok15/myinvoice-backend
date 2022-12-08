package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type CustomerRepository interface {
	GetCustomerDetail(ctx context.Context, customerID int) (*dto.CustomerDetails, error)
}
