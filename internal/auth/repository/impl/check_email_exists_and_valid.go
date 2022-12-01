package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *customerRepository) CheckEmailExistAndValid(ctx context.Context, params *dto.CustomerRequest) (exists, valid bool, err error) {
	getCustomerVerifyStatusSQL, args, err := squirrel.
		Select("cs.is_verified").
		From("customers as c").
		InnerJoin("customer_settings as cs ON cs.customer_id = c.id").
		Where(squirrel.Eq{"c.email": params.Email}).
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
