package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/notifications"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/paymentstatusutils"
)

func (s *invoiceService) ConfirmPayment(ctx context.Context, invoiceID int) error {
	invoice, err := s.repo.GetInvoiceByID(ctx, invoiceID)
	if err != nil {
		return err
	}
	err = s.repo.UpdatePaymentStatus(ctx, invoiceID, paymentstatusutils.Pending)
	if err != nil {
		return err
	}

	err = s.repoNotif.CreateNotificationMerchant(ctx, &dto.CreateNotification{CustomerID: invoice.CustomerID, MerchantID: invoice.MerchantID, InvoiceID: invoiceID, NotificationTitleID: notifications.PaymentDone})
	if err != nil {
		return err
	}

	return nil
}
