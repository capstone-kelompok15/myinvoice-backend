package impl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/emailutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/passwordutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/randomutils"
)

func (s *authService) CustomerRegistration(ctx context.Context, params *dto.CustomerRequest) error {
	exist, valid, err := s.repo.CheckCustomerEmailExistAndValid(ctx, params.Email)
	if err != nil && err != sql.ErrNoRows {
		s.log.Warningln("[CustomerRegistration] Error while checking the existence of the email", err.Error())
		return err
	}

	if exist && valid {
		return customerrors.ErrAccountDuplicated
	}

	if !exist {
		params.Password = passwordutils.HashPassword(params.Password)

		err = s.repo.CustomerRegistration(ctx, params)
		if err != nil {
			s.log.Warningln("[CustomerRegistration] Error while calling the repo function", err.Error())
			return err
		}
	}

	code := randomutils.GenerateNRandomString(4)
	code = strings.ToUpper(code)
	// DONE
	body, err := emailutils.ParseTemplate(emailutils.EmailVerificationCustomer, struct{ Code string }{Code: code})
	if err != nil {
		s.log.Warningln("email error : ", err.Error())
		return err
	}
	mg := s.mailgun.NewMessage(
		s.config.Mailgun.SenderEmail,
		"myInvoice - Your Email Verification Code",
		body,
		params.Email,
	)
	mg.SetHtml(body)
	_, _, err = s.mailgun.Send(ctx, mg)
	if err != nil {
		s.log.Warningln("[CustomerRegistration] Error while send the email:", err.Error())
		return err
	}

	err = s.redis.Set(ctx, fmt.Sprintf("customer-regis:%s", params.Email), code, 5*time.Minute).Err()
	if err != nil {
		s.log.Warningln("[CustomerRegistration] Error while setting code to the redis:", err.Error())
		return err
	}

	return nil
}
