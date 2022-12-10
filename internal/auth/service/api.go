package service

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type AuthService interface {
	CustomerRegistration(ctx context.Context, params *dto.CustomerRequest) error
	CustomerEmailVerification(ctx context.Context, req *dto.CustomerEmailVerification) error
	RefreshEmailVerificationCode(ctx context.Context, email string) error
	CustomerLogin(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerAccessToken, error)
	MerchantRegistration(ctx context.Context, req *dto.MerchantRegisterRequest) error
	LoginAdmin(ctx context.Context, req *dto.AdminLoginRequest) (*dto.AdminLoginResponse, error)
	GenerateNewAccessToken(ctx context.Context, refreshTokens string) (*string, error)
	CustomerResetPasswordRequest(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, req *dto.CustomerResetPassword) error
	AdminEmailVerification(ctx context.Context, req *dto.AdminEmailVerification) error
	RefreshAdminEmailVerificationCode(ctx context.Context, email string) error
	AdminResetPasswordRequest(ctx context.Context, email string) error
	AdminResetPassword(ctx context.Context, req *dto.AdminResetPassword) error
}
