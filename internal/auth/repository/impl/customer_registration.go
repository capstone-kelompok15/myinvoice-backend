package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *customerRepository) CustomerRegistration(ctx context.Context, params *dto.CustomerRequest) error {
	insertCustomerSQL, args, err := squirrel.
		Insert("customers").
		Columns("email", "customer_password").
		Values(params.Email, params.Password).
		ToSql()
	if err != nil {
		r.log.Warningln("[CustomerRegistration] Error while create customers sql from squirrel", err.Error())
		return err
	}

	insertCustomerDetailSQL, _, err := squirrel.
		Insert("customer_details").
		Columns("customer_id", "full_name").
		Values(1, params.FullName).
		ToSql()
	if err != nil {
		r.log.Warningln("[CustomerRegistration] Error while create customer_details sql from squirrel", err.Error())
		return err
	}

	insertCustomerSettingSQL, _, err := squirrel.
		Insert("customer_settings").
		Columns("customer_id", "is_verified", "is_deactivated").
		Values(1, false, true).
		ToSql()
	if err != nil {
		r.log.Warningln("[CustomerRegistration] Error while create customer_details sql from squirrel", err.Error())
		return err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.log.Warningln("[CustomerRegistration] Error while create transaction", err.Error())
		return err
	}
	defer tx.Rollback()

	res, err := tx.ExecContext(ctx, insertCustomerSQL, args...)
	if err != nil {
		r.log.Warningln("[CustomerRegistration] Error while execute ", err.Error())
		return err
	}

	customerID, err := res.LastInsertId()
	if err != nil {
		r.log.Warningln("[CustomerRegistration] Error while getting the last customer ID", err.Error())
		return err
	}

	_, err = tx.ExecContext(ctx, insertCustomerDetailSQL, customerID, params.FullName)
	if err != nil {
		r.log.Warningln("[CustomerRegistration] Error while execute ", err.Error())
		return err
	}

	_, err = tx.ExecContext(ctx, insertCustomerSettingSQL, customerID, false, true)
	if err != nil {
		r.log.Warningln("[CustomerRegistration] Error while execute ", err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		r.log.Warningln("[CustomerRegistration] Error while commit the transaction", err.Error())
		return err
	}

	return nil
}
