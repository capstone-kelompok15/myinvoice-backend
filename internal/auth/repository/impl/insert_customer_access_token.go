package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *authRepository) InsertCustomerAccessToken(ctx context.Context, customerID int, deviceID string) error {
	insertCustomerAccessTokenSQL, args, err := squirrel.
		Insert("customer_tokens").
		Columns("customer_id", "device_id", "is_login").
		Values(customerID, deviceID, true).
		ToSql()
	if err != nil {
		r.log.Warningln("[InsertCustomerAccessToken] Error while creating sql from squirrel", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, insertCustomerAccessTokenSQL, args...)
	if err != nil {
		r.log.Warningln("[InsertCustomerAccessToken] Error while exec the sql", err.Error())
		return err
	}

	return nil
}
