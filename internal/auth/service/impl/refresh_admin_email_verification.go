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
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/randomutils"
)

func (s *authService) RefreshAdminEmailVerificationCode(ctx context.Context, email string) error {
	exist, verif, err := s.repo.CheckAdminEmailExistAndValid(ctx, email)
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

	code := randomutils.GenerateNRandomString(32)
	code = strings.ToUpper(code)
	s.redis.Set(ctx, fmt.Sprintf(dto.MerchantEmailVerifRedisKey, email), code, 5*time.Minute)

	// TODO: Update the callback to the front end callback
	hyperLink := fmt.Sprintf("%s?code=%s&email=%s", "Callback", code, email)
	body, err := emailutils.ParseTemplate(emailutils.EmailVerificationMerchant, struct{ Link string }{Link: hyperLink})
	if err != nil {
		s.log.Warningln("email error : ", err.Error())
		return err
	}
	mg := s.mailgun.NewMessage(
		s.config.Mailgun.SenderEmail,
		"myInvoice - Your New Email Verification Code",
		body,
		hyperLink,
	)
	mg.SetHtml(body)
	_, _, err = s.mailgun.Send(ctx, mg)
	if err != nil {
		s.log.Warningln("[CustomerRegistration] Error while send the email:", err.Error())
		return err
	}

	return nil
}
