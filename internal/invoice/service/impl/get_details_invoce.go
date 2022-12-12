package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (s *invoiceService) GetDetailInvoiceByID(ctx context.Context, req *dto.GetDetailsInvoicesRequest) (*dto.GetInvoiceDetailsByIDResponse, error) {
	invoice, err := s.repo.GetDetailInvoiceByID(ctx, req)
	if err != nil {
		if err != customerrors.ErrRecordNotFound {
			s.log.Warningln("[GetDetailInvoiceByID] Error on running service:", err.Error())
		}
		return nil, err
	}
	return invoice, nil
}
