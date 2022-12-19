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

type customerHandlerSuite struct {
	suite.Suite
	e       *echo.Echo
	handler customerHandler
}

func TestSuiteCustoerHandler(t *testing.T) {
	suite.Run(t, new(customerHandlerSuite))
}

func (suite *customerHandlerSuite) SetupSuite() {
	log := logrus.NewEntry(logrus.New())
	validator, _ := validatorutils.New()

	suite.e = echo.New()
	suite.handler = *NewCustomerHandler(&CustomerHandlerParams{
		Log:       log,
		Validator: validator,
	})
}

type mockCustomerService struct {
	funcGetCustomerDetails
	funcGetAllCustomer
	funcUpdateProfilePicture
	funcUpdateCustomer
	funcGetAllNotificationCustomer
	funcMarkNotifCustomerAsRead
	funcGetSummary
	funcGetUnreadNotifCount
}

type funcGetCustomerDetails func(ctx context.Context, req *dto.CustomerContext) (*dto.CustomerDetails, error)
type funcGetAllCustomer func(ctx context.Context, req *dto.GetAllCustomerRequest) (*[]dto.GetAllCustomerRespond, error)
type funcUpdateProfilePicture func(ctx context.Context, userID *int, newProfilePictureURL *string) (*string, error)
type funcUpdateCustomer func(ctx context.Context, customerID *int, newData *dto.CustomerUpdateRequest) error
type funcGetAllNotificationCustomer func(ctx context.Context, customerID int, req *dto.NotificationRequest) (*[]dto.NotificationRespond, error)
type funcMarkNotifCustomerAsRead func(ctx context.Context, NotifID int, CustomerID int) error
type funcGetSummary func(ctx context.Context, customerID int) (*dto.CustomerSummary, error)
type funcGetUnreadNotifCount func(ctx context.Context, CustomerID int) (int, error)

func (s mockCustomerService) GetCustomerDetails(ctx context.Context, req *dto.CustomerContext) (*dto.CustomerDetails, error) {
	return s.funcGetCustomerDetails(ctx, req)
}

func (s mockCustomerService) GetAllCustomer(ctx context.Context, req *dto.GetAllCustomerRequest) (*[]dto.GetAllCustomerRespond, error) {
	return s.funcGetAllCustomer(ctx, req)
}

func (s mockCustomerService) UpdateProfilePicture(ctx context.Context, userID *int, newProfilePictureURL *string) (*string, error) {
	return s.funcUpdateProfilePicture(ctx, userID, newProfilePictureURL)
}

func (s mockCustomerService) UpdateCustomer(ctx context.Context, customerID *int, newData *dto.CustomerUpdateRequest) error {
	return s.funcUpdateCustomer(ctx, customerID, newData)
}

func (s mockCustomerService) GetAllNotificationCustomer(ctx context.Context, customerID int, req *dto.NotificationRequest) (*[]dto.NotificationRespond, error) {
	return s.funcGetAllNotificationCustomer(ctx, customerID, req)
}

func (s mockCustomerService) MarkNotifCustomerAsRead(ctx context.Context, NotifID int, CustomerID int) error {
	return s.funcMarkNotifCustomerAsRead(ctx, NotifID, CustomerID)
}

func (s mockCustomerService) GetSummary(ctx context.Context, customerID int) (*dto.CustomerSummary, error) {
	return s.funcGetSummary(ctx, customerID)
}

func (s mockCustomerService) GetUnreadNotifCount(ctx context.Context, CustomerID int) (int, error) {
	return s.funcGetUnreadNotifCount(ctx, CustomerID)
}
