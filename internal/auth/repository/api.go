package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type CustomerRepository interface {
	CheckEmailExistAndValid(ctx context.Context, params *dto.CustomerRequest) (exists, valid bool, err error)
	CustomerRegistration(ctx context.Context, params *dto.CustomerRequest) error
}
