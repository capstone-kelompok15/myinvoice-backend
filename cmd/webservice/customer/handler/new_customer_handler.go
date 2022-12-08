package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/internal/customer/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/sirupsen/logrus"
)

type customerHandler struct {
	service   service.CustomerService
	log       *logrus.Entry
	validator *validatorutils.Validator
}

type CustomerHandlerParams struct {
	Service   service.CustomerService
	Log       *logrus.Entry
	Validator *validatorutils.Validator
}

func NewCustomerHandler(params *CustomerHandlerParams) *customerHandler {
	return &customerHandler{
		service:   params.Service,
		log:       params.Log,
		validator: params.Validator,
	}
}
