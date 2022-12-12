package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type InvoiceRepository interface {
	CreateInvoice(ctx context.Context, merchantID int, req *dto.CreateInvoiceRequest) error
}
