package impl

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/randomutils"
)

func (s *customerService) CustomerResetPasswordRequest(ctx context.Context, email string) error {
	exist, valid, err := s.repo.CheckCustomerEmailExistAndValid(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return customerrors.ErrRecordNotFound
		}
		s.log.Warningln("[CustomerResetPasswordRequest] Error while checking the existence of the email")
		return err
	}

	if !valid || !exist {
		return customerrors.ErrUnauthorized
	}

	code := randomutils.GenerateNRandomString(128)
	frontEndCallback := fmt.Sprintf("%s?code=%s&email=%s", s.config.FrontEndURL, code, email)

	mg := s.mailgun.NewMessage(
		s.config.Mailgun.SenderEmail,
		"myInvoice - Your Reset Password Request",
		// TODO: Need to change to the html template
		frontEndCallback,
		email,
	)

	_, _, err = s.mailgun.Send(ctx, mg)
	if err != nil {
		s.log.Warningln("[CustomerResetPasswordRequest] Error while send the email:", err.Error())
		return err
	}

	err = s.redis.Set(ctx, fmt.Sprintf("customer-reset-password:%s", email), code, 10*time.Minute).Err()
	if err != nil {
		s.log.Warningln("[CustomerResetPasswordRequest] Error while setting code to the redis:", err.Error())
		return err
	}

	return nil
}
