package impl

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *customerRepository) CheckEmailExistAndValid(ctx context.Context, params *dto.CustomerRequest) (exists, valid bool, err error) {
	subQuerySQL, arg, err := squirrel.
		Select("id").
		From("customers").
		Where(squirrel.Eq{"email": params.Email}).
		ToSql()
	if err != nil {
		return false, false, err
	}

	getCustomerVerifyStatusSQL, _, err := squirrel.
		Select("is_verified").
		From("customer_settings").
		Where(fmt.Sprintf("customer_id = (%s)", subQuerySQL)).
		ToSql()
	if err != nil {
		return false, false, err
	}

	var isVerified bool

	row := r.db.QueryRowContext(ctx, getCustomerVerifyStatusSQL, arg...)
	err = row.Scan(&isVerified)
	if err != nil {
		return false, false, err
	}

	if !isVerified {
		return true, false, nil
	}

	return true, true, nil
}
