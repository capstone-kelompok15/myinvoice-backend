package impl

import (
	"context"
	"fmt"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/emailutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/notifications"
)

func (s *invoiceService) CreateInvoice(ctx context.Context, merchantID int, req *dto.CreateInvoiceRequest) error {
	fullName, email, err := s.repo.GetCustomerByID(ctx, req.CustomerID)
	if err != nil {
		if err != customerrors.ErrRecordNotFound {
			s.log.Warningln("[CreateInvoice] Error on getting customer info by ID:", err.Error())
		}
		return err
	}

	invoiceID, err := s.repo.CreateInvoice(ctx, merchantID, req)
	if err != nil {
		s.log.Warningln("[CreateInvoice] error on service:", err.Error())
		return err
	}

	content := fmt.Sprintf("Hi %s, you have new invoice with id #%d", *fullName, invoiceID)
	body, err := emailutils.ParseTemplate(emailutils.EmailNotifNewInvoice, struct{ Content string }{Content: content})
	if err != nil {
		s.log.Warningln("email error : ", err.Error())
		return err
	}
	mg := s.mailgun.NewMessage(
		s.config.Mailgun.SenderEmail,
		"myInvoice - You Have New Invoice",
		// TODO: Need to change to the html template
		body,
		*email,
	)
	mg.SetHtml(body)
	_, _, err = s.mailgun.Send(ctx, mg)
	if err != nil {
		s.log.Warningln("[CreateInvoice] Error while send the email:", err.Error())
		return err
	}
	// send notif to customer
	err = s.repoNotif.CreateNotificationCustomer(ctx, &dto.CreateNotification{CustomerID: req.CustomerID, MerchantID: merchantID, InvoiceID: invoiceID, NotificationTitleID: notifications.NewBillIssued})
	if err != nil {
		return err
	}

	return nil
}
