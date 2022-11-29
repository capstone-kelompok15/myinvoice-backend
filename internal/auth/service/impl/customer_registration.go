package impl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/passwordutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/randomutils"
)

func (s customerService) CustomerRegistration(ctx context.Context, params *dto.CustomerRequest) error {
	exist, valid, err := s.repo.CheckEmailExistAndValid(ctx, params)
	if err != nil && err != sql.ErrNoRows {
		s.log.Warningln("[CustomerRegistration] Error while checking the existence of the email")
		return err
	}

	if exist && valid {
		return customerrors.ErrAccountDuplicated
	}

	if !exist {
		params.Password = passwordutils.HashPassword(params.Password)

		err = s.repo.CustomerRegistration(ctx, params)
		if err != nil {
			s.log.Warningln("[CustomerRegistration] Error while calling the repo function")
			return err
		}
	}

	code := randomutils.GenerateNRandomString(4)
	code = strings.ToUpper(code)
	s.redis.Set(ctx, fmt.Sprintf("regis:%s", params.Email), code, 5*time.Minute)

	mg := s.mailgun.NewMessage(
		s.config.Mailgun.SenderEmail,
		"myInvoice - Your Email Verification Code",
		// TODO: Need to change to the html template
		code,
		params.Email,
	)

	_, _, err = s.mailgun.Send(ctx, mg)
	if err != nil {
		s.log.Warningln("[CustomerRegistration] Error while send the email:", err.Error())
		return err
	}

	return nil
}
