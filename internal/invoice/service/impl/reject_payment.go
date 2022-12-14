package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/notifications"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/paymentstatusutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/websocketutils"
)

func (s *invoiceService) RejectPayment(ctx context.Context, invoiceID int, message string) (*websocketutils.Message, error) {
	invoice, err := s.repo.GetInvoiceByID(ctx, invoiceID)
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdatePaymentStatus(ctx, invoiceID, paymentstatusutils.Failed)
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdateMessage(ctx, invoiceID, message)
	if err != nil {
		return nil, err
	}

	err = s.repoNotif.CreateNotificationCustomer(ctx, &dto.CreateNotification{CustomerID: invoice.CustomerID, MerchantID: invoice.MerchantID, InvoiceID: invoiceID, NotificationTitleID: notifications.PaymentFailed})
	if err != nil {
		return nil, err
	}

	messageWebsocket := websocketutils.NewWebSocketMessage(&websocketutils.MessageParams{
		Content:        "payment rejected",
		InvoiceID:      invoiceID,
		SendToCustomer: invoice.CustomerID,
		SendToMerchant: invoice.MerchantID,
		PaymentTypeID:  paymentstatusutils.Failed,
		PaymentType:    "Rejected",
	})

	return messageWebsocket, nil
}
