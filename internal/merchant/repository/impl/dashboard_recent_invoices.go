package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *merchantRepository) GetDashboardRecentInvoices(ctx context.Context, merchantID int) (*[]dto.RecentInvoiceMerchantDashboard, error) {
	recentInvoicesSQL, args, err := squirrel.
		Select("cd.full_name as customer_name", "SUM(id.price) as price", "i.id as invoice_id", "i.due_at as invoice_expired_date").
		From("invoices as i").
		InnerJoin("invoice_details as id on i.id = id.invoice_id").
		InnerJoin("customer_details as cd on i.customer_id = cd.customer_id").
		GroupBy("id.invoice_id").
		Limit(7).
		Where(squirrel.Eq{"i.merchant_id": merchantID}).
		OrderBy("i.due_at desc").
		ToSql()
	if err != nil {
		return nil, err
	}

	recentInvoices := []dto.RecentInvoiceMerchantDashboard{}

	err = r.db.SelectContext(ctx, &recentInvoices, recentInvoicesSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			recentInvoices = []dto.RecentInvoiceMerchantDashboard{}
		} else {
			r.log.Warningln("[GetDashboardRecentInvoices] Error on execute query", err.Error())
			return nil, err
		}
	}

	return &recentInvoices, nil
}
