package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *authRepository) InvalidCustomerAccessToken(ctx context.Context, customerID int) error {
	invalidateAccessTokenSQL, args, err := squirrel.
		Update("customer_tokens").
		Set("is_login", false).
		Where(squirrel.Eq{"customer_id": customerID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[InvalidCustomerAccessToken] Error while creating sql from squirrel", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, invalidateAccessTokenSQL, args...)
	if err != nil {
		r.log.Warningln("[InvalidCustomerAccessToken] Error while exec the query", err.Error())
		return err
	}

	return nil
}
