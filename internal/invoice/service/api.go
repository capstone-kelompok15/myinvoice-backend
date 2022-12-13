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
	UploadPayment(ctx context.Context, customerID int, invoiceID int, filePath string) error
	ConfirmPayment(ctx context.Context, invoiceID int) error
	AcceptPayment(ctx context.Context, invoiceID int) error
	RejectPayment(ctx context.Context, invoiceID int, message string) error
	GetReport(ctx context.Context, params *dto.ReportParams) (*dto.ReportResponse, error)
}
