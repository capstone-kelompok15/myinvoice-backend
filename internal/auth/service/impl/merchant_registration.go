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

func (s *customerService) MerchantRegistration(ctx context.Context, req *dto.MerchantRegisterRequest) error {
	exist, valid, err := s.repo.CheckAdminEmailExistAndValid(ctx, req)
	if err != nil && err != sql.ErrNoRows {
		s.log.Warningln("[CustomerRegistration] Error while checking the existence of the email", err.Error())
		return err
	}

	if exist && valid {
		return customerrors.ErrAccountDuplicated
	}

	if !exist {
		req.Password = passwordutils.HashPassword(req.Password)

		err = s.repo.MerchantRegistration(ctx, req)
		if err != nil {
			s.log.Warningln("[CustomerRegistration] Error while calling the repo function", err.Error())
			return err
		}
	}

	code := randomutils.GenerateNRandomString(32)
	code = strings.ToUpper(code)

	// TODO: Update the callback to the front end callback
	hyperLink := fmt.Sprintf("%s?code=%s&email=%s", "Callback", code, req.Email)

	mg := s.mailgun.NewMessage(
		s.config.Mailgun.SenderEmail,
		"myInvoice - Your Email Verification Code",
		// TODO: Need to change to the html template
		hyperLink,
		req.Email,
	)

	_, _, err = s.mailgun.Send(ctx, mg)
	if err != nil {
		s.log.Warningln("[CustomerRegistration] Error while send the email:", err.Error())
		return err
	}

	err = s.redis.Set(ctx, fmt.Sprintf("merchant-regis:%s", req.Email), code, 10*time.Minute).Err()
	if err != nil {
		s.log.Warningln("[CustomerRegistration] Error while setting code to the redis:", err.Error())
		return err
	}

	return nil
}
