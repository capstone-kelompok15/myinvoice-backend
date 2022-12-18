package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type InvoiceRepository interface {
	GetCustomerByID(ctx context.Context, customerID int) (fullName, email *string, err error)
	CreateInvoice(ctx context.Context, merchantID int, req *dto.CreateInvoiceRequest) (int, error)
	GetAllInvoice(ctx context.Context, req *dto.GetAllInvoicesParam) (*[]dto.GetInvoiceResponse, int, error)
	GetDetailInvoiceByID(ctx context.Context, req *dto.GetDetailsInvoicesRequest) (*dto.GetInvoiceDetailsByIDResponse, error)
	GetCustomers(ctx context.Context, req *dto.GetMerchantCustomerList) (*[]dto.BriefCustomer, int, error)
	UpdatePaymentStatus(ctx context.Context, invoiceID int, paymentStatusID int) error
	ValidateInvoiceID(ctx context.Context, customerID int, invoiceID int, invoiceDetailID *int) error
	UploadPayment(ctx context.Context, invoiceID int, uploadedURL string) error
	GetMerchantProfile(ctx context.Context, invoiceID int) (*dto.MerchantBriefDate, error)
	GetInvoiceByID(ctx context.Context, invoiceID int) (*dto.GetInvoiceByID, error)
	GetReport(ctx context.Context, params *dto.ReportParams) (*dto.ReportResponse, error)
	UpdateMessage(ctx context.Context, invoiceID int, message string) error
	UpdateMerchantBankID(ctx context.Context, invoiceID int, merchantBankID int) error
	GetPaymentStatusList(ctx context.Context) (*[]dto.PaymentStatus, error)
	DeleteInvoice(ctx context.Context, req *dto.DeleteInvoice) error
	DeleteDetailInvoice(ctx context.Context, req *dto.DeleteInvoice) error
}
