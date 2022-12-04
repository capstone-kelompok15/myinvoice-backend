package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *customerRepository) GetAdminContextByID(ctx context.Context, adminID int) (*dto.AdminContext, error) {
	loginSQL, args, err := squirrel.
		Select("a.id as id", "a.merchant_id as merchant_id", "m.merchant_name as merchant_name").
		From("admins as a").
		InnerJoin("merchants as m ON m.id = a.id").
		Where(squirrel.Eq{"a.id": adminID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetAdminContextByID] Error while creating sql from squirrel", err.Error())
		return nil, err
	}

	var adminCtx dto.AdminContext
	err = r.db.Get(&adminCtx, loginSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[GetAdminContextByID] Error while exec the query", err.Error())
		return nil, err
	}
	return nil, nil
}
