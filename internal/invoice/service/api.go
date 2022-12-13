package service

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type InvoiceService interface {
	CreateInvoice(ctx context.Context, merchantID int, req *dto.CreateInvoiceRequest) error
	GetAllInvoice(ctx context.Context, req *dto.GetAllInvoicesParam) (*[]dto.GetInvoiceResponse, int, error)
	GetDetailInvoiceByID(ctx context.Context, req *dto.GetDetailsInvoicesRequest) (*dto.GetInvoiceDetailsByIDResponse, error)
	GetCustomers(ctx context.Context, req *dto.GetMerchantCustomerList) (*[]dto.BriefCustomer, int, error)
	ConfirmPayment(ctx context.Context, invoiceID int) error
}
