package impl

import (
	"context"
)

func (s *invoiceService) UpdatePaymentMethod(ctx context.Context, invoiceID, merchantBankID int) error {
	_, err := s.repo.GetInvoiceByID(ctx, invoiceID)
	if err != nil {
		return err
	}
	err = s.repo.UpdateMerchantBankID(ctx, invoiceID, merchantBankID)
	if err != nil {
		return err
	}

	return nil
}
