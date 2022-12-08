package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type MiddlewareRepository interface {
	ValidateCustomer(ctx context.Context, deviceID *string, userID *int) (*dto.CustomerContext, error)
}
