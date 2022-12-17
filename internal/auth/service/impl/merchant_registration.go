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

func (s *authService) MerchantRegistration(ctx context.Context, req *dto.MerchantRegisterRequest) error {
	exist, valid, err := s.repo.CheckAdminEmailExistAndValid(ctx, req.Email)
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
			if err == customerrors.ErrMerchantNameDuplicated || err == customerrors.ErrUsernameDuplicated {
				s.log.Warningln("[CustomerRegistration] Error while calling the repo function", err.Error())
			}
			return err
		}
	}

	code := randomutils.GenerateNRandomString(32)
	code = strings.ToUpper(code)

	// TODO: Update the callback to the front end callback
	hyperLink := fmt.Sprintf("%s?code=%s&email=%s", s.config.FrontEndURL, code, req.Email)
	body, err := emailutils.ParseTemplate(emailutils.EmailVerificationMerchant, struct{ Link string }{Link: hyperLink})
	if err != nil {
		s.log.Warningln("email error : ", err.Error())
		return err
	}
	mg := s.mailgun.NewMessage(
		s.config.Mailgun.SenderEmail,
		"myInvoice - Your Email Verification Code",
		body,
		req.Email,
	)
	mg.SetHtml(body)

	_, _, err = s.mailgun.Send(ctx, mg)
	if err != nil {
		s.log.Warningln("[CustomerRegistration] Error while send the email:", err.Error())
		return err
	}

	err = s.redis.Set(ctx, fmt.Sprintf(dto.MerchantEmailVerifRedisKey, req.Email), code, 10*time.Minute).Err()
	if err != nil {
		s.log.Warningln("[CustomerRegistration] Error while setting code to the redis:", err.Error())
		return err
	}

	return nil
}
