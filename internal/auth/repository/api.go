package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type CustomerRepository interface {
	CheckCustomerEmailExistAndValid(ctx context.Context, params *dto.CustomerRequest) (exists, valid bool, err error)
	CustomerRegistration(ctx context.Context, params *dto.CustomerRequest) error
	CustomerEmailVerification(ctx context.Context, req *dto.CustomerEmailVerification) error
	AuthorizeCustomerLogin(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerContext, error)
	InvalidCustomerAccessToken(ctx context.Context, customerID int) error
	InsertCustomerAccessToken(ctx context.Context, customerID int, deviceID string) error
	CheckAdminEmailExistAndValid(ctx context.Context, params *dto.MerchantRegisterRequest) (exists, valid bool, err error)
	MerchantRegistration(ctx context.Context, req *dto.MerchantRegisterRequest) error
	LoginAdmin(ctx context.Context, req *dto.AdminLoginRequest) (*dto.AdminContext, error)
	InsertRefreshToken(ctx context.Context, req *dto.CustomerRefreshToken) error
}
