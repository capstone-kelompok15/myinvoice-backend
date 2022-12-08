package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *authRepository) InsertRefreshToken(ctx context.Context, req *dto.AdminRefreshToken) error {
	insertRefreshTokenSQL, args, err := squirrel.
		Insert("refresh_tokens").
		Columns("admin_id", "token", "is_valid", "expired_date").
		Values(req.AdminID, req.Token, req.IsValid, req.ExpirationDate).
		ToSql()
	if err != nil {
		r.log.Warningln("[InsertRefreshToken] Error while creating sql from squirrel", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, insertRefreshTokenSQL, args...)
	if err != nil {
		r.log.Warningln("[InsertRefreshToken] Error while execution sql", err.Error())
		return err
	}

	return nil
}
