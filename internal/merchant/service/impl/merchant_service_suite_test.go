package impl

import (
	"context"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/config"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type merchantServiceSuite struct {
	suite.Suite
	service merchantService
}

func TestSuiteMerchantService(t *testing.T) {
	suite.Run(t, new(merchantServiceSuite))
}

func (suite *merchantServiceSuite) SetupSuite() {
	configTest, err := config.GetConfig("./../../../../.env")
	if err != nil {
		panic(err)
	}

	cld, err := config.GetCloudinaryConn(&configTest.Cloudinary)
	if err != nil {
		panic(err)
	}

	suite.service = *NewMerchantService(&MerchantServiceParams{
		Log:        logrus.NewEntry(logrus.New()),
		Config:     configTest,
		Cloudinary: cld,
	})
}

type mockNotificationRepository struct {
	funcGetTitleID
	funcCheckNotifCustomerExist
	funcCheckNotifMerchantExist
	funcGetAllNotificationCustomer
	funcGetAllNotificationMerchant
	funcCreateNotificationCustomer
	funcCreateNotificationMerchant
	funcMarkNotifCustomerAsRead
	funcMarkNotifMerchantAsRead
	funcGetUnreadNotifCountMerchant
	funcGetUnreadNotifCountCustomer
}

type funcGetTitleID func(title string) (int, error)
type funcCheckNotifCustomerExist func(ID int) error
type funcCheckNotifMerchantExist func(ID int) error
type funcGetAllNotificationCustomer func(ctx context.Context, customerID int, req *dto.NotificationRequest) (*[]dto.CustomerNotificationDB, error)
type funcGetAllNotificationMerchant func(ctx context.Context, merchantID int, req *dto.NotificationRequest) (*[]dto.MerchantNotificationDB, error)
type funcCreateNotificationCustomer func(ctx context.Context, req *dto.CreateNotification) error
type funcCreateNotificationMerchant func(ctx context.Context, req *dto.CreateNotification) error
type funcMarkNotifCustomerAsRead func(ctx context.Context, NotifID int, CustomerID int) error
type funcMarkNotifMerchantAsRead func(ctx context.Context, NotifID int, MerchantID int) error
type funcGetUnreadNotifCountMerchant func(ctx context.Context, MerchantID int) (int, error)
type funcGetUnreadNotifCountCustomer func(ctx context.Context, CustomerID int) (int, error)

func (m mockNotificationRepository) GetTitleID(title string) (int, error) {
	return m.funcGetTitleID(title)
}

func (m mockNotificationRepository) CheckNotifCustomerExist(ID int) error {
	return m.funcCheckNotifCustomerExist(ID)
}

func (m mockNotificationRepository) CheckNotifMerchantExist(ID int) error {
	return m.funcCheckNotifMerchantExist(ID)
}

func (m mockNotificationRepository) GetAllNotificationCustomer(ctx context.Context, customerID int, req *dto.NotificationRequest) (*[]dto.CustomerNotificationDB, error) {
	return m.funcGetAllNotificationCustomer(ctx, customerID, req)
}

func (m mockNotificationRepository) GetAllNotificationMerchant(ctx context.Context, merchantID int, req *dto.NotificationRequest) (*[]dto.MerchantNotificationDB, error) {
	return m.funcGetAllNotificationMerchant(ctx, merchantID, req)
}

func (m mockNotificationRepository) CreateNotificationCustomer(ctx context.Context, req *dto.CreateNotification) error {
	return m.funcCreateNotificationCustomer(ctx, req)
}

func (m mockNotificationRepository) CreateNotificationMerchant(ctx context.Context, req *dto.CreateNotification) error {
	return m.funcCreateNotificationMerchant(ctx, req)
}

func (m mockNotificationRepository) MarkNotifCustomerAsRead(ctx context.Context, NotifID int, CustomerID int) error {
	return m.funcMarkNotifCustomerAsRead(ctx, NotifID, CustomerID)
}

func (m mockNotificationRepository) MarkNotifMerchantAsRead(ctx context.Context, NotifID int, MerchantID int) error {
	return m.funcMarkNotifMerchantAsRead(ctx, NotifID, MerchantID)
}

func (m mockNotificationRepository) GetUnreadNotifCountMerchant(ctx context.Context, MerchantID int) (int, error) {
	return m.funcGetUnreadNotifCountMerchant(ctx, MerchantID)
}

func (m mockNotificationRepository) GetUnreadNotifCountCustomer(ctx context.Context, CustomerID int) (int, error) {
	return m.funcGetUnreadNotifCountCustomer(ctx, CustomerID)
}

type mockMerchantRepository struct {
	funcGetDashboardInvoiceOverview
	funcGetDashboardRecentInvoices
	funcGetDashboardRecentPayments
	funcGetMerchantBank
	funcGetMerchantProfile
	funcUpdateMerchantBank
	funcValidateMerchantBank
	funcCreateMerchantBank
	funcUpdateProfilePicture
	funcUpdateMerchantProfile
}

type funcGetDashboardInvoiceOverview func(ctx context.Context, merchantID int) (*dto.OverviewMerchantDashboard, error)
type funcGetDashboardRecentInvoices func(ctx context.Context, merchantID int) (*[]dto.RecentInvoiceMerchantDashboard, error)
type funcGetDashboardRecentPayments func(ctx context.Context, merchantID int) (*[]dto.RecentPaymentMerchantDashboard, error)
type funcGetMerchantBank func(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error)
type funcGetMerchantProfile func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error)
type funcUpdateMerchantBank func(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error
type funcValidateMerchantBank func(ctx context.Context, merchantID int, merchantBankID int) error
type funcCreateMerchantBank func(ctx context.Context, merchantID int, req *dto.MerchantBankData) error
type funcUpdateProfilePicture func(ctx context.Context, merchantID *int, newProfilePictureURL *string) error
type funcUpdateMerchantProfile func(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error

func (r mockMerchantRepository) GetDashboardInvoiceOverview(ctx context.Context, merchantID int) (*dto.OverviewMerchantDashboard, error) {
	return r.funcGetDashboardInvoiceOverview(ctx, merchantID)
}

func (r mockMerchantRepository) GetDashboardRecentInvoices(ctx context.Context, merchantID int) (*[]dto.RecentInvoiceMerchantDashboard, error) {
	return r.funcGetDashboardRecentInvoices(ctx, merchantID)
}

func (r mockMerchantRepository) GetDashboardRecentPayments(ctx context.Context, merchantID int) (*[]dto.RecentPaymentMerchantDashboard, error) {
	return r.funcGetDashboardRecentPayments(ctx, merchantID)
}

func (r mockMerchantRepository) GetMerchantBank(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error) {
	return r.funcGetMerchantBank(ctx, req)
}

func (r mockMerchantRepository) GetMerchantProfile(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
	return r.funcGetMerchantProfile(ctx, merchantID)
}

func (r mockMerchantRepository) UpdateMerchantBank(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error {
	return r.funcUpdateMerchantBank(ctx, req)
}

func (r mockMerchantRepository) ValidateMerchantBank(ctx context.Context, merchantID int, merchantBankID int) error {
	return r.funcValidateMerchantBank(ctx, merchantID, merchantBankID)
}

func (r mockMerchantRepository) CreateMerchantBank(ctx context.Context, merchantID int, req *dto.MerchantBankData) error {
	return r.funcCreateMerchantBank(ctx, merchantID, req)
}

func (r mockMerchantRepository) UpdateProfilePicture(ctx context.Context, merchantID *int, newProfilePictureURL *string) error {
	return r.funcUpdateProfilePicture(ctx, merchantID, newProfilePictureURL)
}

func (r mockMerchantRepository) UpdateMerchantProfile(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
	return r.funcUpdateMerchantProfile(ctx, merchantID, req)
}
