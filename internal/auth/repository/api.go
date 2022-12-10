package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type AuthRepository interface {
	CheckCustomerEmailExistAndValid(ctx context.Context, email string) (exists, valid bool, err error)
	CustomerRegistration(ctx context.Context, params *dto.CustomerRequest) error
	CustomerEmailVerification(ctx context.Context, req *dto.CustomerEmailVerification) error
	AuthorizeCustomerLogin(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerContext, error)
	InvalidCustomerAccessToken(ctx context.Context, customerID int) error
	InsertCustomerAccessToken(ctx context.Context, customerID int, deviceID string) error
	CheckAdminEmailExistAndValid(ctx context.Context, email string) (exists, valid bool, err error)
	MerchantRegistration(ctx context.Context, req *dto.MerchantRegisterRequest) error
	LoginAdmin(ctx context.Context, req *dto.AdminLoginRequest) (*dto.AdminContext, error)
	InsertRefreshToken(ctx context.Context, req *dto.AdminRefreshToken) error
	GetRefreshToken(ctx context.Context, refreshToken string) (*dto.AdminRefreshToken, error)
	InvalidateRefreshToken(ctx context.Context, refreshToken *dto.AdminRefreshToken) error
	GetAdminContextByID(ctx context.Context, adminID int) (*dto.AdminContext, error)
	AdminEmailVerification(ctx context.Context, req *dto.AdminEmailVerification) error
	CustomerResetPassword(ctx context.Context, req *dto.CustomerResetPassword) error
}
