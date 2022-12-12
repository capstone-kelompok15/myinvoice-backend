package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (s *invoiceService) CreateInvoice(ctx context.Context, merchantID int, req *dto.CreateInvoiceRequest) error {
	err := s.repo.CreateInvoice(ctx, merchantID, req)
	if err != nil {
		s.log.Warningln("[CreateInvoice] error on service:", err.Error())
		return err
	}

	return nil
}
