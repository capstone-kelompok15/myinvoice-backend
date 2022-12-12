package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type InvoiceRepository interface {
	GetCustomerByID(ctx context.Context, customerID int) (fullName, email *string, err error)
	CreateInvoice(ctx context.Context, merchantID int, req *dto.CreateInvoiceRequest) error
	GetAllInvoice(ctx context.Context, req *dto.GetAllInvoicesParam) (*[]dto.GetInvoiceResponse, error)
}
