package emailutils

import (
	"bytes"
	"html/template"
)

var (
	EmailResetPasswordReq      = "assets/email-reset-password-req.html"
	EmailVerificationCustomer  = "assets/email-verification-customer.html"
	EmailVerificationMerchant  = "assets/email-verification-merchant.html"
	EmailNotifNewInvoice       = "assets/email-notif-new-invoice.html"
	EmailNotifResetPassSuccess = "assets/email-notif-reset-password-success.html"
	EmailNotifCustomerHasPaid  = "assets/email-notif-customer-has-paid.html"
)

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}
