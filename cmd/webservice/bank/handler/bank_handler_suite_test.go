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

type bankHandlerSuite struct {
	suite.Suite
	e       *echo.Echo
	handler bankHandler
}

func TestSuiteBankHandler(t *testing.T) {
	suite.Run(t, new(bankHandlerSuite))
}

func (suite *bankHandlerSuite) SetupSuite() {
	log := logrus.NewEntry(logrus.New())
	validator, _ := validatorutils.New()

	suite.e = echo.New()
	suite.handler = *NewBankHandler(&BankHandler{
		Log:       log,
		Validator: validator,
	})
}

type mockBankService struct {
	funcGetAllBank
}

type funcGetAllBank func(ctx context.Context) (*[]dto.BankResponse, error)

func (s mockBankService) GetAllBank(ctx context.Context) (*[]dto.BankResponse, error) {
	return s.funcGetAllBank(ctx)
}
