package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *customerRepository) GetAllCustomer(ctx context.Context, req *dto.GetAllCustomerRequest) (*[]dto.GetAllCustomerRespond, error) {
	startIndex := (req.Page - 1) * req.Limit
	getAllCustomerSQL, args, err := squirrel.
		Select("c.id as id", "c.email as email", "c.created_at as created_at", "c.updated_at as updated_at", "cd.full_name as full_name", "cd.display_profile_url as display_profile_url", "cd.address as address").
		From("customers as c").
		InnerJoin("customer_details as cd on c.id = cd.customer_id").
		InnerJoin("customer_settings as cs on c.id = cs.customer_id").
		Where(squirrel.NotEq{"cs.is_verified": false}).
		Where(squirrel.Like{"cd.full_name": "%" + req.Name + "%"}).
		Where(squirrel.Like{"c.email": "%" + req.Email + "%"}).
		OrderBy("c.id DESC").
		Offset(startIndex).Limit(req.Limit).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetAllCustomer] Error while creating sql from squirrel", err.Error())
		return nil, err
	}

	var customers []dto.GetAllCustomerRespond
	err = r.db.Select(&customers, getAllCustomerSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[GetAllCustomer] Error while exec the query", err.Error())
		return nil, err
	}

	return &customers, nil
}
