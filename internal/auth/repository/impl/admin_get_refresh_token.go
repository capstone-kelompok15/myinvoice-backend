package impl

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *authRepository) GetRefreshToken(ctx context.Context, refreshToken string) (*dto.AdminRefreshToken, error) {
	now := time.Now().String()
	getRefreshTokenSQL, args, err := squirrel.
		Select("*").
		From("refresh_tokens as rt ").
		Where(squirrel.Eq{"is_valid": true}).
		Where(squirrel.GtOrEq{"expired_at": now}).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetRefreshToken] Error while sql from squirrel:", err.Error())
		return nil, err
	}

	var refreshTokenRes dto.AdminRefreshToken

	err = r.db.GetContext(ctx, &refreshToken, getRefreshTokenSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrUnauthorized
		}
		r.log.Warningln("[GetRefreshToken] Error while querying refresh token:", err.Error())
		return nil, err
	}

	return &refreshTokenRes, nil
}
