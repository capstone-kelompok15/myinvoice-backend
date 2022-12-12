package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *invoiceRepository) GetCustomerByID(ctx context.Context, customerID int) (fullName, email *string, err error) {
	getCustomerByIDSQL, args, err := squirrel.
		Select("c.email AS email, cd.full_name AS full_name").
		From("customers AS c").
		InnerJoin("customer_details AS cd ON cd.customer_id = c.id").
		Where(squirrel.Eq{"c.id": customerID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetCustomerByID] Failed on build sql:", err.Error())
		return nil, nil, err
	}

	customer := struct {
		Email    string `db:"email"`
		FullName string `db:"full_name"`
	}{}

	err = r.db.GetContext(ctx, &customer, getCustomerByIDSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[GetCustomerByID] Failed on executing query:", err.Error())
		return nil, nil, err
	}

	return &customer.FullName, &customer.Email, nil
}
