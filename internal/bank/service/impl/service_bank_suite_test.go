package impl

import (
	"context"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type bankServiceSuite struct {
	suite.Suite
	service bankService
}

func TestSuiteBankService(t *testing.T) {
	suite.Run(t, new(bankServiceSuite))
}

func (suite *bankServiceSuite) SetupSuite() {
	suite.service = *NewBankService(&BankServiceParams{
		Repo: nil,
		Log:  logrus.NewEntry(logrus.New()),
	})
}

type mockBankRepository struct {
	funcGetAllBank
}

type funcGetAllBank func(ctx context.Context) (*[]dto.BankResponse, error)

func (s mockBankRepository) GetAllBank(ctx context.Context) (*[]dto.BankResponse, error) {
	return s.funcGetAllBank(ctx)
}
