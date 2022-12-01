package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type CustomerRepository interface {
	CheckEmailExistAndValid(ctx context.Context, params *dto.CustomerRequest) (exists, valid bool, err error)
	CustomerRegistration(ctx context.Context, params *dto.CustomerRequest) error
	CustomerEmailVerification(ctx context.Context, req *dto.CustomerEmailVerification) error
	AuthorizeCustomerLogin(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerContext, error)
	InvalidCustomerAccessToken(ctx context.Context, customerID int) error
	InsertCustomerAccessToken(ctx context.Context, customerID int, deviceID string) error
}
