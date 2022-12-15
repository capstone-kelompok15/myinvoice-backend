package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/dateutils"
)

func (r *invoiceRepository) GetAllInvoice(ctx context.Context, req *dto.GetAllInvoicesParam) (*[]dto.GetInvoiceResponse, int, error) {
	getAllInvoiceBuilder := squirrel.StatementBuilder

	if req.CustomerID != 0 {
		getAllInvoiceBuilder = getAllInvoiceBuilder.Where(squirrel.Eq{"i.customer_id": req.CustomerID})
	} else if req.MerchantID != 0 {
		getAllInvoiceBuilder = getAllInvoiceBuilder.Where(squirrel.Eq{"i.merchant_id": req.MerchantID})
		if req.MerchantFilterCustomerID != 0 {
			getAllInvoiceBuilder = getAllInvoiceBuilder.Where(squirrel.Eq{"i.customer_id": req.MerchantFilterCustomerID})
		}
	}

	if req.DateFilter != nil {
		getAllInvoiceBuilder = getAllInvoiceBuilder.Where(squirrel.GtOrEq{"i.due_at": dateutils.TimeToDateStr(req.DateFilter.StartDate)})
		getAllInvoiceBuilder = getAllInvoiceBuilder.Where(squirrel.LtOrEq{"i.due_at": dateutils.TimeToDateStr(req.DateFilter.EndDate)})
	}

	if req.PaymentStatusID != 0 {
		getAllInvoiceBuilder = getAllInvoiceBuilder.Where(squirrel.Eq{"i.payment_status_id": req.PaymentStatusID})
	}

	getAllInvoiceSQL, args, err := getAllInvoiceBuilder.
		Select(
			"i.id AS invoice_id", "i.merchant_id AS merchant_id", "m.merchant_name AS merchant_name",
			"cd.full_name AS customer_name", "ps.status_name AS payment_status_name",
			"i.payment_status_id AS payment_status_id", "pt.payment_type_name AS payment_type_name",
			"i.due_at AS due_at", "i.created_at AS created_at", "i.updated_at AS updated_at",
			"SUM(id.price * id.quantity) AS total_price",
		).
		From("invoices AS i").
		InnerJoin("invoice_details AS id ON id.invoice_id = i.id").
		InnerJoin("payment_statuses AS ps ON ps.id = i.payment_status_id").
		InnerJoin("customer_details AS cd ON cd.customer_id = i.customer_id").
		InnerJoin("merchants AS m ON m.id = i.merchant_id").
		LeftJoin("payment_types AS pt ON pt.id = i.payment_type_id").
		GroupBy("id.invoice_id").
		OrderBy("i.due_at DESC").
		Limit(uint64(req.PaginationFilter.Limit)).
		Offset(uint64(req.PaginationFilter.Offset)).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetAllInvoice] Failed on build get all invoice SQL:", err.Error())
		return nil, 0, err
	}

	countInvoicesSQL, args1, err := getAllInvoiceBuilder.
		Select("COUNT(i.id)").
		From("invoices AS i").
		GroupBy("i.id").
		ToSql()
	if err != nil {
		r.log.Warningln("[GetAllInvoice] Failed on build get count invoice SQL:", err.Error())
		return nil, 0, err
	}

	invoices := []dto.GetInvoiceResponse{}
	err = r.db.SelectContext(ctx, &invoices, getAllInvoiceSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return &invoices, 0, nil
		}
		r.log.Warningln("[GetAllInvoice] Failed on executing get all invoice sql:", err.Error())
		return nil, 0, err
	}

	var count int
	err = r.db.GetContext(ctx, &count, countInvoicesSQL, args1...)
	if err != nil {
		if err == sql.ErrNoRows {
			return &invoices, 0, nil
		}
		r.log.Warningln("[GetAllInvoice] Failed on executing count invoice sql:", err.Error())
		return nil, 0, err
	}

	return &invoices, count, nil
}
