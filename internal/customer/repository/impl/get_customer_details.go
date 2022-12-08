package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *customerRepository) GetCustomerDetail(ctx context.Context, customerID int) (*dto.CustomerDetails, error) {
	getCustomerDetailSQL, args, err := squirrel.
		Select("c.id as id", "c.email as email", "cd.full_name as full_name", "cd.diplay_profile_url as display_profile_url").
		From("customers as c").
		InnerJoin("customer_details as cd on c.id = cd.customer_id").
		Where(squirrel.Eq{"c.id": customerID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetCustomerDetail] Error while creating sql from squirrel", err.Error())
		return nil, err
	}

	var customerDetails dto.CustomerDetails
	err = r.db.Get(&customerDetails, getCustomerDetailSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[GetCustomerDetail] Error while exec the query", err.Error())
		return nil, err
	}

	return &customerDetails, nil
}
