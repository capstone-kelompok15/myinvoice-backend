package impl

import (
	"context"
	"fmt"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/emailutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/passwordutils"
	"github.com/go-redis/redis/v8"
)

func (s *authService) AdminResetPassword(ctx context.Context, req *dto.AdminResetPassword) error {
	var err error

	key := fmt.Sprintf(dto.MerchantResetPasswordRedisKey, req.Email)
	res := s.redis.Get(ctx, key)
	err = res.Err()
	if err != nil {
		if err == redis.Nil {
			return customerrors.ErrNotFound
		}
		s.log.Warningln("[AdminResetPassword] Error while setting code to the redis:", err.Error())
		return err
	}

	if string(res.Val()) != req.Code {
		return customerrors.ErrNotFound
	}

	err = s.redis.Del(ctx, key).Err()
	if err != nil {
		s.log.Warningln("[AdminResetPassword] Error while destroy the key:", err.Error())
		return err
	}

	req.Password = passwordutils.HashPassword(req.Password)
	err = s.repo.AdminResetPassword(ctx, req)
	if err != nil {
		s.log.Warningln("[AdminResetPassword] Error while update the password:", err.Error())
		return err
	}
	body, err := emailutils.ParseTemplate(emailutils.EmailResetPasswordReq, struct{}{})
	if err != nil {
		s.log.Warningln("email error : ", err.Error())
		return err
	}
	mg := s.mailgun.NewMessage(
		s.config.Mailgun.SenderEmail,
		"myInvoice - Your Password Was Reset",
		// TODO: Need to change to the html template
		body,
		req.Email,
	)
	mg.SetHtml(body)
	_, _, err = s.mailgun.Send(ctx, mg)
	if err != nil {
		s.log.Warningln("[AdminResetPassword] Error while send the email:", err.Error())
		return err
	}

	return nil
}
