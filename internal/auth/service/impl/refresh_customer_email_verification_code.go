package impl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/randomutils"
)

func (s *authService) RefreshEmailVerificationCode(ctx context.Context, email string) error {
	exist, verif, err := s.repo.CheckCustomerEmailExistAndValid(ctx, email)
	if err != nil {
		if err != sql.ErrNoRows {
			return customerrors.ErrRecordNotFound
		}
		s.log.Warningln("[CustomerRegistration] Error while checking the existence of the email")
		return err
	}

	if exist && verif {
		return customerrors.ErrBadRequest
	}

	code := randomutils.GenerateNRandomString(4)
	code = strings.ToUpper(code)
	s.redis.Set(ctx, fmt.Sprintf("customer-regis:%s", email), code, 5*time.Minute)

	mg := s.mailgun.NewMessage(
		s.config.Mailgun.SenderEmail,
		"myInvoice - Your New Email Verification Code",
		// TODO: Need to change to the html template
		code,
		email,
	)

	_, _, err = s.mailgun.Send(ctx, mg)
	if err != nil {
		s.log.Warningln("[CustomerRegistration] Error while send the email:", err.Error())
		return err
	}

	return nil
}
