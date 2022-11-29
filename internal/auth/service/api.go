package service

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type CustomerService interface {
	CustomerRegistration(ctx context.Context, params *dto.CustomerRequest) error
}
