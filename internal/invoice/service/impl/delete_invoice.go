package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (s *invoiceService) DeleteInvoice(ctx context.Context, req *dto.DeleteInvoice) error {
	var invoiceDetailID *int

	invoiceDetailID = &req.InvoiceDetailID
	if req.InvoiceDetailID == 0 {
		invoiceDetailID = nil
	}

	err := s.repo.ValidateInvoiceID(ctx, req.CustomerID, req.InvoiceID, invoiceDetailID)
	if err != nil {
		if err != customerrors.ErrRecordNotFound {
			s.log.Warningln("[DeleteInvoice] Error on validate invoice ID", err.Error())
		}
		return err
	}

	if req.InvoiceDetailID == 0 {
		err := s.repo.DeleteInvoice(ctx, req)
		if err != nil {
			s.log.Warningln("[DeleteInvoice] Error on running the repository:", err.Error())
			return err
		}
	} else {
		err := s.repo.DeleteDetailInvoice(ctx, req)
		if err != nil {
			s.log.Warningln("[DeleteInvoice] Error on running the repository:", err.Error())
			return err
		}
	}
	return nil
}
