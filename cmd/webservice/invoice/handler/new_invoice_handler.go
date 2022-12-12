package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/internal/invoice/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/sirupsen/logrus"
)

type invoiceHandler struct {
	service   service.InvoiceService
	log       *logrus.Entry
	validator *validatorutils.Validator
}

type InvoiceHandlerParams struct {
	Service   service.InvoiceService
	Log       *logrus.Entry
	Validator *validatorutils.Validator
}

func NewInvoiceHandler(params *InvoiceHandlerParams) *invoiceHandler {
	return &invoiceHandler{
		service:   params.Service,
		log:       params.Log,
		validator: params.Validator,
	}
}
