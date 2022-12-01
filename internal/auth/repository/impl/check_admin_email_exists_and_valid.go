package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *customerRepository) CheckAdminEmailExistAndValid(ctx context.Context, params *dto.MerchantRegisterRequest) (exists, valid bool, err error) {
	getCustomerVerifyStatusSQL, args, err := squirrel.
		Select("is_verified").
		From("admins").
		Where(squirrel.Eq{"email": params.Email}).
		ToSql()
	if err != nil {
		return false, false, err
	}

	var isVerified bool
	row := r.db.QueryRowContext(ctx, getCustomerVerifyStatusSQL, args...)
	err = row.Scan(&isVerified)
	if err != nil {
		return false, false, err
	}

	if !isVerified {
		return true, false, nil
	}

	return true, true, nil
}
