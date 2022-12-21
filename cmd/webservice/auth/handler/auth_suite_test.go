package handler

import (
	"context"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type authHandlerSuite struct {
	suite.Suite
	e       *echo.Echo
	handler authHandler
}

func TestSuiteAuthHandler(t *testing.T) {
	suite.Run(t, new(authHandlerSuite))
}

func (suite *authHandlerSuite) SetupSuite() {
	log := logrus.NewEntry(logrus.New())
	validator, _ := validatorutils.New()

	suite.e = echo.New()
	suite.handler = *NewAuthHandler(&AuthHandlerParams{
		Log:       log,
		Validator: validator,
	})
}

type mockAuthService struct {
	funcCustomerRegistration
	funcCustomerEmailVerification
	funcRefreshEmailVerificationCode
	funcCustomerLogin
	funcMerchantRegistration
	funcLoginAdmin
	funcGenerateNewAccessToken
	funcCustomerResetPasswordRequest
	funcResetPassword
	funcAdminEmailVerification
	funcRefreshAdminEmailVerificationCode
	funcAdminResetPasswordRequest
	funcAdminResetPassword
}

type funcCustomerRegistration func(ctx context.Context, params *dto.CustomerRequest) error
type funcCustomerEmailVerification func(ctx context.Context, req *dto.CustomerEmailVerification) error
type funcRefreshEmailVerificationCode func(ctx context.Context, email string) error
type funcCustomerLogin func(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerAccessToken, error)
type funcMerchantRegistration func(ctx context.Context, req *dto.MerchantRegisterRequest) error
type funcLoginAdmin func(ctx context.Context, req *dto.AdminLoginRequest) (*dto.AdminLoginResponse, error)
type funcGenerateNewAccessToken func(ctx context.Context, refreshTokens string) (*string, error)
type funcCustomerResetPasswordRequest func(ctx context.Context, email string) error
type funcResetPassword func(ctx context.Context, req *dto.CustomerResetPassword) error
type funcAdminEmailVerification func(ctx context.Context, req *dto.AdminEmailVerification) error
type funcRefreshAdminEmailVerificationCode func(ctx context.Context, email string) error
type funcAdminResetPasswordRequest func(ctx context.Context, email string) error
type funcAdminResetPassword func(ctx context.Context, req *dto.AdminResetPassword) error

func (m mockAuthService) CustomerRegistration(ctx context.Context, params *dto.CustomerRequest) error {
	return m.funcCustomerRegistration(ctx, params)
}

func (m mockAuthService) CustomerEmailVerification(ctx context.Context, req *dto.CustomerEmailVerification) error {
	return m.funcCustomerEmailVerification(ctx, req)
}

func (m mockAuthService) RefreshEmailVerificationCode(ctx context.Context, email string) error {
	return m.funcRefreshEmailVerificationCode(ctx, email)
}

func (m mockAuthService) CustomerLogin(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerAccessToken, error) {
	return m.funcCustomerLogin(ctx, req)
}

func (m mockAuthService) MerchantRegistration(ctx context.Context, req *dto.MerchantRegisterRequest) error {
	return m.funcMerchantRegistration(ctx, req)
}

func (m mockAuthService) LoginAdmin(ctx context.Context, req *dto.AdminLoginRequest) (*dto.AdminLoginResponse, error) {
	return m.funcLoginAdmin(ctx, req)
}

func (m mockAuthService) GenerateNewAccessToken(ctx context.Context, refreshTokens string) (*string, error) {
	return m.funcGenerateNewAccessToken(ctx, refreshTokens)
}

func (m mockAuthService) CustomerResetPasswordRequest(ctx context.Context, email string) error {
	return m.funcCustomerResetPasswordRequest(ctx, email)
}

func (m mockAuthService) ResetPassword(ctx context.Context, req *dto.CustomerResetPassword) error {
	return m.funcResetPassword(ctx, req)
}

func (m mockAuthService) AdminEmailVerification(ctx context.Context, req *dto.AdminEmailVerification) error {
	return m.funcAdminEmailVerification(ctx, req)
}

func (m mockAuthService) RefreshAdminEmailVerificationCode(ctx context.Context, email string) error {
	return m.funcRefreshAdminEmailVerificationCode(ctx, email)
}

func (m mockAuthService) AdminResetPasswordRequest(ctx context.Context, email string) error {
	return m.funcAdminResetPasswordRequest(ctx, email)
}

func (m mockAuthService) AdminResetPassword(ctx context.Context, req *dto.AdminResetPassword) error {
	return m.funcAdminResetPassword(ctx, req)
}
