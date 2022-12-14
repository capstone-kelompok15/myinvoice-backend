package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/internal/invoice/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/websocketutils"
	"github.com/sirupsen/logrus"
)

type invoiceHandler struct {
	service       service.InvoiceService
	log           *logrus.Entry
	validator     *validatorutils.Validator
	websocketPool *websocketutils.Pool
}

type InvoiceHandlerParams struct {
	Service       service.InvoiceService
	Log           *logrus.Entry
	Validator     *validatorutils.Validator
	WebsocketPool *websocketutils.Pool
}

func NewInvoiceHandler(params *InvoiceHandlerParams) *invoiceHandler {
	return &invoiceHandler{
		service:       params.Service,
		log:           params.Log,
		validator:     params.Validator,
		websocketPool: params.WebsocketPool,
	}
}
