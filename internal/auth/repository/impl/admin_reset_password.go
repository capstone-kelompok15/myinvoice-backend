package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *authRepository) AdminResetPassword(ctx context.Context, req *dto.AdminResetPassword) error {
	adminResetPasswordSQL, args, err := squirrel.
		Update("admins").
		Set("admin_password", req.Password).
		Where(squirrel.Eq{"email": req.Email}).
		ToSql()
	if err != nil {
		r.log.Warningln("[AdminResetPassword] Error while create customers sql from squirrel", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, adminResetPasswordSQL, args...)
	if err != nil {
		r.log.Warningln("[AdminResetPassword] Error while create customers sql from squirrel", err.Error())
		return err
	}

	return nil
}
