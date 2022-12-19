package impl

import (
	"context"
	"fmt"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	cloudinary "github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/cloudinaryutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/emailutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/notifications"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/paymentstatusutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/websocketutils"
)

func (s *invoiceService) UploadPayment(ctx context.Context, customerID int, invoiceID int, filePath string) (*websocketutils.Message, error) {
	err := s.repo.ValidateInvoiceID(ctx, customerID, invoiceID, nil)
	if err != nil {
		if err != customerrors.ErrRecordNotFound {
			s.log.Warningln("[UploadPayment] Failed on running validate repository:", err.Error())
		}
		return nil, err
	}

	imageURL, err := cloudinary.UploadFile(cloudinary.UploadFileParams{
		Ctx:      ctx,
		Cld:      s.cloudinary,
		Filename: filePath,
	})
	if err != nil {
		s.log.Warningln("[UploadPayment] Failed on upload the file:", err.Error())
		return nil, err
	}

	err = s.repo.UploadPayment(ctx, invoiceID, *imageURL)
	if err != nil {
		s.log.Warningln("[UploadPayment] Failed on inserting repository:", err.Error())
		return nil, err
	}

	merchantBrief, err := s.repo.GetMerchantProfile(ctx, invoiceID)
	if err != nil {
		s.log.Warningln("[UploadPayment] Failed on getting merchant brief:", err.Error())
		return nil, err
	}

	err = s.repoNotif.CreateNotificationMerchant(ctx, &dto.CreateNotification{
		CustomerID:          customerID,
		MerchantID:          merchantBrief.MerchantID,
		InvoiceID:           invoiceID,
		NotificationTitleID: notifications.PaymentDone,
	})
	if err != nil {
		return nil, err
	}

	customerName, _, err := s.repo.GetCustomerByID(ctx, customerID)

	content := fmt.Sprintf("%s has paid the invoice with id %d", *customerName, invoiceID)
	body, err := emailutils.ParseTemplate(emailutils.EmailNotifCustomerHasPaid, struct {
		Image   string
		Content string
	}{Image: *imageURL, Content: content})

	mg := s.mailgun.NewMessage(
		s.config.Mailgun.SenderEmail,
		"myInvoice - You Have New Payment",
		body,
		merchantBrief.Email,
	)
	mg.SetHtml(body)

	_, _, err = s.mailgun.Send(ctx, mg)
	if err != nil {
		s.log.Warningln("[UploadPayment] Error while send the email:", err.Error())
		return nil, err
	}

	message := websocketutils.NewWebSocketMessage(&websocketutils.MessageParams{
		Content:        "payment approval uploaded",
		InvoiceID:      invoiceID,
		SendToCustomer: customerID,
		SendToMerchant: merchantBrief.MerchantID,
		PaymentTypeID:  paymentstatusutils.Pending,
		PaymentType:    "Pending",
	})

	return message, err
}
