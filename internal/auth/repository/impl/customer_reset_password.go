package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *authRepository) CustomerResetPassword(ctx context.Context, req *dto.CustomerResetPassword) error {
	customerResetPasswordSQL, args, err := squirrel.
		Update("customers").
		Set("customer_password", req.Password).
		Where(squirrel.Eq{"email": req.Email}).
		ToSql()
	if err != nil {
		r.log.Warningln("[CustomerResetPassword] Error while create customers sql from squirrel", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, customerResetPasswordSQL, args...)
	if err != nil {
		r.log.Warningln("[CustomerResetPassword] Error while create customers sql from squirrel", err.Error())
		return err
	}

	return nil
}
