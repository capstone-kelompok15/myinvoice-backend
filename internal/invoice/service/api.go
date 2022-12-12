package service

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type InvoiceService interface {
	CreateInvoice(ctx context.Context, merchantID int, req *dto.CreateInvoiceRequest) error
}
