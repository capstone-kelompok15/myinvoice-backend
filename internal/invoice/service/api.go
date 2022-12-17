package service

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/websocketutils"
)

type InvoiceService interface {
	CreateInvoice(ctx context.Context, merchantID int, req *dto.CreateInvoiceRequest) error
	GetAllInvoice(ctx context.Context, req *dto.GetAllInvoicesParam) (*[]dto.GetInvoiceResponse, int, error)
	GetDetailInvoiceByID(ctx context.Context, req *dto.GetDetailsInvoicesRequest) (*dto.GetInvoiceDetailsByIDResponse, error)
	GetCustomers(ctx context.Context, req *dto.GetMerchantCustomerList) (*[]dto.BriefCustomer, int, error)
	UploadPayment(ctx context.Context, customerID int, invoiceID int, filePath string) (*websocketutils.Message, error)
	ConfirmPayment(ctx context.Context, invoiceID int) (*websocketutils.Message, error)
	AcceptPayment(ctx context.Context, invoiceID int) (*websocketutils.Message, error)
	RejectPayment(ctx context.Context, invoiceID int, message string) (*websocketutils.Message, error)
	GetReport(ctx context.Context, params *dto.ReportParams) (*dto.ReportResponse, error)
	UpdatePaymentMethod(ctx context.Context, invoiceID int, merchantBankID int) error
	GetPaymentStatusList(ctx context.Context) (*[]dto.PaymentStatus, error)
	GeneratePDF(ctx context.Context, req *dto.GetDetailsInvoicesRequest, downloadBase string) (*string, error)
}
