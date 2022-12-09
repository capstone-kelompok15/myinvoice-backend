package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *merchantRepository) GetDashboardRecentPayments(ctx context.Context, merchantID int) (*[]dto.RecentPaymentMerchantDashboard, error) {
	recentPaymentSQL, args, err := squirrel.
		Select(
			"cd.full_name as customer_name",
			"SUM(id.price) as price",
			"i.id as invoice_id",
			"i.due_at as invoice_expired_date",
			"pt.payment_type_name as payment_type").
		From("invoices as i").
		InnerJoin("invoice_details as id on i.id = id.invoice_id").
		InnerJoin("customer_details as cd on i.customer_id = cd.customer_id").
		InnerJoin("payment_types as pt on i.payment_type_id = pt.id").
		GroupBy("id.invoice_id").
		Limit(7).
		Where(squirrel.Eq{"i.merchant_id": merchantID}).
		OrderBy("i.due_at desc").
		ToSql()
	if err != nil {
		return nil, err
	}

	recentPayment := []dto.RecentPaymentMerchantDashboard{}

	err = r.db.SelectContext(ctx, &recentPayment, recentPaymentSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			recentPayment = []dto.RecentPaymentMerchantDashboard{}
		} else {
			r.log.Warningln("[GetDashboardRecentPayments] Error on execute query:", err.Error())
			return nil, err
		}
	}

	return &recentPayment, nil
}
