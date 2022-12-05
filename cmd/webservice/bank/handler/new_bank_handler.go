package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/internal/bank/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/sirupsen/logrus"
)

type bankHandler struct {
	service   service.BankService
	log       *logrus.Entry
	validator *validatorutils.Validator
}

type BankHandler struct {
	Service   service.BankService
	Log       *logrus.Entry
	Validator *validatorutils.Validator
}

func NewBankHandler(params *BankHandler) *bankHandler {
	return &bankHandler{
		service:   params.Service,
		log:       params.Log,
		validator: params.Validator,
	}
}
