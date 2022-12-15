package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/notifications"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/paymentstatusutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/websocketutils"
)

func (s *invoiceService) AcceptPayment(ctx context.Context, invoiceID int) (*websocketutils.Message, error) {
	invoice, err := s.repo.GetInvoiceByID(ctx, invoiceID)
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdatePaymentStatus(ctx, invoiceID, paymentstatusutils.Paid)
	if err != nil {
		return nil, err
	}

	err = s.repoNotif.CreateNotificationCustomer(ctx, &dto.CreateNotification{CustomerID: invoice.CustomerID, MerchantID: invoice.MerchantID, InvoiceID: invoiceID, NotificationTitleID: notifications.PaymentSuccess})
	if err != nil {
		return nil, err
	}

	message := websocketutils.NewWebSocketMessage(&websocketutils.MessageParams{
		Content:        "payment accepted",
		InvoiceID:      invoiceID,
		SendToCustomer: invoice.CustomerID,
		SendToMerchant: invoice.MerchantID,
		PaymentTypeID:  paymentstatusutils.Paid,
		PaymentType:    "Paid",
	})

	return message, nil
}
