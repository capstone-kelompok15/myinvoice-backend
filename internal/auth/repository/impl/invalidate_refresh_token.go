package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *authRepository) InvalidateRefreshToken(ctx context.Context, refreshToken *dto.AdminRefreshToken) error {
	getRefreshTokenSQL, args, err := squirrel.
		Update("refresh_tokens").
		Set("is_valid", false).
		Where(squirrel.GtOrEq{"id": refreshToken.ID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[InvalidateRefreshToken] Error while create sql from squirrel:", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, getRefreshTokenSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return customerrors.ErrUnauthorized
		}
		r.log.Warningln("[InvalidateRefreshToken] Error while update refresh token:", err.Error())
		return err
	}
	return nil
}
