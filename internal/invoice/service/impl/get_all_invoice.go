package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (s *invoiceService) GetAllInvoice(ctx context.Context, req *dto.GetAllInvoicesParam) (*[]dto.GetInvoiceResponse, int, error) {
	invoices, count, err := s.repo.GetAllInvoice(ctx, req)
	if err != nil {
		s.log.Warningln("[GetAllInvoice] Failed on running the repository")
		return nil, 0, err
	}
	return invoices, count, nil
}
