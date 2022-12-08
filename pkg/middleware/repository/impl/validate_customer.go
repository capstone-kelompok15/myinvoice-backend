package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *middlewareRepository) ValidateCustomer(ctx context.Context, deviceID *string, userID *int) (*dto.CustomerContext, error) {
	queryCustomerContextSQL, args, err := squirrel.
		Select("ct.customer_id as id", "cd.full_name as full_name").
		From("customer_tokens as ct").
		InnerJoin("customer_details as cd ON cd.customer_id = cd.customer_id").
		Where(squirrel.Eq{"ct.device_id": deviceID}).
		Where(squirrel.Eq{"ct.customer_id": userID}).
		Where(squirrel.Eq{"ct.is_login": true}).
		ToSql()
	if err != nil {
		r.log.Warningln("[ValidateCustomer] Error while creating query:", err.Error())
		return nil, err
	}

	var customerCtx dto.CustomerContext
	err = r.db.GetContext(ctx, &customerCtx, queryCustomerContextSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[ValidateCustomer] Error while execute the query:", err.Error())
		return nil, err
	}

	return &customerCtx, nil
}
