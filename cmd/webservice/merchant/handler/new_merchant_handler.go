package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/internal/merchant/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/sirupsen/logrus"
)

type merchantHandler struct {
	service   service.MerchantService
	log       *logrus.Entry
	validator *validatorutils.Validator
}

type MerchantHandler struct {
	Service   service.MerchantService
	Log       *logrus.Entry
	Validator *validatorutils.Validator
}

func NewMerchantHandler(params *MerchantHandler) *merchantHandler {
	return &merchantHandler{
		service:   params.Service,
		log:       params.Log,
		validator: params.Validator,
	}
}
