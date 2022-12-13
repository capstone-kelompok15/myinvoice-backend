package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *invoiceRepository) GetCustomers(ctx context.Context, req *dto.GetMerchantCustomerList) (*[]dto.BriefCustomer, int, error) {
	getCustomerSQL, args, err := squirrel.
		Select("i.customer_id AS id", "c.email AS email", "cd.full_name AS full_name").
		From("invoices AS i").
		InnerJoin("customers AS c ON c.id = i.customer_id").
		InnerJoin("customer_details AS cd ON cd.customer_id = i.customer_id").
		Where(squirrel.Eq{"i.merchant_id": req.MerchantID}).
		GroupBy("i.customer_id").
		ToSql()
	if err != nil {
		r.log.Warningln("[GetCustomer] Failed on build sql:", err.Error())
		return nil, 0, err
	}

	countCustomerSQL, args1, err := squirrel.
		Select("COUNT(i.id)").
		From("invoices AS i").
		Where(squirrel.Eq{"i.merchant_id": req.MerchantID}).
		GroupBy("i.customer_id").
		ToSql()
	if err != nil {
		r.log.Warningln("[GetCustomer] Failed on build sql for counting:", err.Error())
		return nil, 0, err
	}

	customers := []dto.BriefCustomer{}

	err = r.db.SelectContext(ctx, &customers, getCustomerSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return &customers, 0, nil
		}
		r.log.Warningln("[GetCustomer] Error on query:", err.Error())
		return nil, 0, err
	}

	var count int
	err = r.db.GetContext(ctx, &count, countCustomerSQL, args1...)
	if err != nil {
		if err == sql.ErrNoRows {
			return &customers, 0, nil
		}
		r.log.Warningln("[GetCustomer] Error on counting:", err.Error())
		return nil, 0, err
	}

	return &customers, count, nil
}
