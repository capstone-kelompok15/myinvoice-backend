package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *invoiceRepository) GetDetailInvoiceByID(ctx context.Context, req *dto.GetDetailsInvoicesRequest) (*dto.GetInvoiceDetailsByIDResponse, error) {
	getAllInvoiceSQL, args, err := squirrel.
		Select(
			"i.id AS invoice_id", "i.merchant_id AS merchant_id",
			"i.customer_id AS customer_id", "cd.full_name AS customer_name", "cd.address AS customer_address",
			"i.payment_status_id AS payment_status_id", "ps.status_name AS payment_status_name",
			"i.payment_type_id AS payment_type_id", "pt.payment_type_name AS payment_type_name",
			"i.merchant_bank_id AS merchant_bank_id", "SUM(id.price * id.quantity) AS total_price",
			"COUNT(i.id) AS product_quantity", "i.note AS note", "i.message AS message",
			"i.due_at AS due_at", "i.created_at AS created_at", "i.updated_at AS updated_at",
		).
		From("invoices AS i").
		InnerJoin("invoice_details AS id ON id.invoice_id = i.id").
		InnerJoin("payment_statuses AS ps ON ps.id = i.payment_status_id").
		InnerJoin("customer_details AS cd ON cd.customer_id = i.customer_id").
		LeftJoin("payment_types AS pt ON pt.id = i.payment_type_id").
		Where(squirrel.Eq{"i.id": req.InvoiceID}).
		GroupBy("id.invoice_id").
		OrderBy("created_at DESC").
		ToSql()
	if err != nil {
		r.log.Warningln("[GetAllInvoice] Failed on build get all invoice SQL:", err.Error())
		return nil, err
	}

	getInvoiceDetailsSQL, _, err := squirrel.
		Select(
			"id AS invoice_detail_id", "product", "quantity",
			"price", "created_at", "updated_at",
		).
		From("invoice_details").
		Where(squirrel.Eq{"invoice_id": "?"}).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetAllInvoice] Failed on build get detail invoice SQL:", err.Error())
		return nil, err
	}

	prepInvoiceDetail, err := r.db.PreparexContext(ctx, getInvoiceDetailsSQL)
	if err != nil {
		r.log.Warningln("[GetAllInvoice] Failed on prepared detail invoice statement:", err.Error())
		return nil, err
	}
	defer prepInvoiceDetail.Close()

	invoices := dto.GetInvoiceDetailsByIDResponse{}

	err = r.db.GetContext(ctx, &invoices, getAllInvoiceSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[GetAllInvoice] Failed on executing get all invoice sql:", err.Error())
		return nil, err
	}

	invoices.InvoiceDetail = []dto.GetInvoiceDetail{}

	err = r.db.SelectContext(ctx, &invoices.InvoiceDetail, getInvoiceDetailsSQL, invoices.InvoiceID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[GetAllInvoice] Failed on executing get detail invoice sql:", err.Error())
		return nil, err
	}

	return &invoices, nil
}
