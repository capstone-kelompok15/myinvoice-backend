package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/notifications"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/paymentstatusutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/websocketutils"
)

func (s *invoiceService) ConfirmPayment(ctx context.Context, invoiceID int) (*websocketutils.Message, error) {
	invoice, err := s.repo.GetInvoiceByID(ctx, invoiceID)
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdatePaymentStatus(ctx, invoiceID, paymentstatusutils.Pending)
	if err != nil {
		return nil, err
	}

	err = s.repoNotif.CreateNotificationMerchant(ctx, &dto.CreateNotification{CustomerID: invoice.CustomerID, MerchantID: invoice.MerchantID, InvoiceID: invoiceID, NotificationTitleID: notifications.PaymentDone})
	if err != nil {
		return nil, err
	}

	message := websocketutils.NewWebSocketMessage(&websocketutils.MessageParams{
		Content:        "payment confirmed",
		InvoiceID:      invoiceID,
		SendToCustomer: invoice.CustomerID,
		SendToMerchant: invoice.MerchantID,
		PaymentTypeID:  paymentstatusutils.Pending,
		PaymentType:    "Pending",
	})

	return message, nil
}
