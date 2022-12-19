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

type merchantHandlerSuite struct {
	suite.Suite
	e       *echo.Echo
	handler merchantHandler
}

func TestSuiteMerchantHandler(t *testing.T) {
	suite.Run(t, new(merchantHandlerSuite))
}

func (suite *merchantHandlerSuite) SetupSuite() {
	log := logrus.NewEntry(logrus.New())
	validator, _ := validatorutils.New()

	suite.e = echo.New()
	suite.handler = *NewMerchantHandler(&MerchantHandler{
		Log:       log,
		Validator: validator,
	})
}

type mockMerchantService struct {
	funcGetAllNotificationMerchant
	funcMarkNotifMerchantAsRead
	funcGetDashboard
	funcGetMerchantBank
	funcUpdateMerchantBank
	funcCreateMerchantBank
	funcUpdateProfilePicture
	funcGetMerchantProfile
	funcGetUnreadNotifCount
	funcUpdateMerchantProfile
}

type funcGetAllNotificationMerchant func(ctx context.Context, merchantID int, req *dto.NotificationRequest) (*[]dto.NotificationRespond, error)
type funcMarkNotifMerchantAsRead func(ctx context.Context, NotifID int, MerchantID int) error
type funcGetDashboard func(ctx context.Context, merchantID int) (*dto.MerchantDashboard, error)
type funcGetMerchantBank func(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error)
type funcUpdateMerchantBank func(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error
type funcCreateMerchantBank func(ctx context.Context, merchantID int, req *dto.MerchantBankData) error
type funcUpdateProfilePicture func(ctx context.Context, userID *int, newProfilePictureURL *string) (*string, error)
type funcGetMerchantProfile func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error)
type funcGetUnreadNotifCount func(ctx context.Context, MerchantID int) (int, error)
type funcUpdateMerchantProfile func(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error

func (s mockMerchantService) GetAllNotificationMerchant(ctx context.Context, merchantID int, req *dto.NotificationRequest) (*[]dto.NotificationRespond, error) {
	return s.funcGetAllNotificationMerchant(ctx, merchantID, req)
}

func (s mockMerchantService) MarkNotifMerchantAsRead(ctx context.Context, NotifID int, MerchantID int) error {
	return s.funcMarkNotifMerchantAsRead(ctx, NotifID, MerchantID)
}

func (s mockMerchantService) GetDashboard(ctx context.Context, merchantID int) (*dto.MerchantDashboard, error) {
	return s.funcGetDashboard(ctx, merchantID)
}

func (s mockMerchantService) GetMerchantBank(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error) {
	return s.funcGetMerchantBank(ctx, req)
}

func (s mockMerchantService) UpdateMerchantBank(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error {
	return s.funcUpdateMerchantBank(ctx, req)
}

func (s mockMerchantService) CreateMerchantBank(ctx context.Context, merchantID int, req *dto.MerchantBankData) error {
	return s.funcCreateMerchantBank(ctx, merchantID, req)
}

func (s mockMerchantService) UpdateProfilePicture(ctx context.Context, userID *int, newProfilePictureURL *string) (*string, error) {
	return s.funcUpdateProfilePicture(ctx, userID, newProfilePictureURL)
}

func (s mockMerchantService) GetMerchantProfile(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
	return s.funcGetMerchantProfile(ctx, merchantID)
}

func (s mockMerchantService) GetUnreadNotifCount(ctx context.Context, MerchantID int) (int, error) {
	return s.funcGetUnreadNotifCount(ctx, MerchantID)
}

func (s mockMerchantService) UpdateMerchantProfile(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
	return s.funcUpdateMerchantProfile(ctx, merchantID, req)
}
