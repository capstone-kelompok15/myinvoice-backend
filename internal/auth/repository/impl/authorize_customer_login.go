package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *customerRepository) AuthorizeCustomerLogin(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerContext, error) {
	loginSQL, args, err := squirrel.
		Select("c.id as id", "cd.full_name as full_name").
		From("customers as c").
		Join("customer_details as cd on cd.customer_id = c.id").
		Where(squirrel.Eq{"c.email": req.Email}).
		Where(squirrel.Eq{"c.customer_password": req.Password}).
		ToSql()
	if err != nil {
		r.log.Warningln("[AuthorizeCustomerLogin] Error while creating sql from squirrel", err.Error())
		return nil, err
	}

	var customerContext dto.CustomerContext
	err = r.db.Get(&customerContext, loginSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrUnauthorized
		}
		r.log.Warningln("[AuthorizeCustomerLogin] Error while getting customer ID", err.Error())
		return nil, err
	}

	return &customerContext, nil
}
