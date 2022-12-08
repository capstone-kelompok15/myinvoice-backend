package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *authRepository) CheckCustomerEmailExistAndValid(ctx context.Context, email string) (exists, valid bool, err error) {
	getCustomerVerifyStatusSQL, args, err := squirrel.
		Select("cs.is_verified").
		From("customers as c").
		InnerJoin("customer_settings as cs ON cs.customer_id = c.id").
		Where(squirrel.Eq{"c.email": email}).
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
