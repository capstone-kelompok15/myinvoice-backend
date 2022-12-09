package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *customerRepository) UpdateCustomer(ctx context.Context, userID *int, newData *dto.CustomerUpdateRequest) error {
	updateSQL, args, err := squirrel.
		Update("customer_details").
		Set("full_name", newData.FullName).
		Set("address", newData.Address).
		Where(squirrel.Eq{"customer_id": *userID}).
		ToSql()
	if err != nil {
		if err == sql.ErrNoRows {
			return customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[UpdateCustomer] Error while creating sql from squirrel", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, updateSQL, args...)
	if err != nil {
		r.log.Warningln("[UpdateCustomer] Error while executing query", err.Error())
		return err
	}

	return nil
}
