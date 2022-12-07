package impl

import (
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	"github.com/capstone-kelompok15/myinvoice-backend/internal/bank/repository"
	"github.com/sirupsen/logrus"
)

type bankService struct {
	repo   repository.BankRepository
	log    *logrus.Entry
	config *config.Config
}

type BankServiceParams struct {
	Repo   repository.BankRepository
	Log    *logrus.Entry
	Config *config.Config
}

func NewBankService(params *BankServiceParams) *bankService {
	return &bankService{
		repo:   params.Repo,
		log:    params.Log,
		config: params.Config,
	}
}
