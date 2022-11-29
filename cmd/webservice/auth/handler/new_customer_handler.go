package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/internal/auth/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/sirupsen/logrus"
)

type customerHandler struct {
	service   service.CustomerService
	log       *logrus.Entry
	validator *validatorutils.Validator
}

type CustomerHandler struct {
	Service   service.CustomerService
	Log       *logrus.Entry
	Validator *validatorutils.Validator
}

func NewCustomerHandler(params *CustomerHandler) *customerHandler {
	return &customerHandler{
		service:   params.Service,
		log:       params.Log,
		validator: params.Validator,
	}
}
