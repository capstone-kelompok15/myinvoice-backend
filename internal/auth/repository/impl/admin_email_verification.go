package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *authRepository) AdminEmailVerification(ctx context.Context, req *dto.AdminEmailVerification) error {
	updateVerifiedStatusSQL, args, err := squirrel.
		Update("admins").
		Set("is_verified", true).
		Where(squirrel.Eq{"email": req.Email}).
		ToSql()
	if err != nil {
		r.log.Warningln("[CustomerEmailVerification] Error while create update sql from squirrel", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, updateVerifiedStatusSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return customerrors.ErrRecordNotFound
		}
		return err
	}

	return nil
}
