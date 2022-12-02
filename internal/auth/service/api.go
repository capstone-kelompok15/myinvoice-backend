package service

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type CustomerService interface {
	CustomerRegistration(ctx context.Context, params *dto.CustomerRequest) error
	CustomerEmailVerification(ctx context.Context, req *dto.CustomerEmailVerification) error
	RefreshEmailVerificationCode(ctx context.Context, email string) error
	CustomerLogin(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerAccessToken, error)
	MerchantRegistration(ctx context.Context, req *dto.MerchantRegisterRequest) error
	LoginAdmin(ctx context.Context, req *dto.AdminLoginRequest) (*dto.AdminLoginResponse, error)
}
