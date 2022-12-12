package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *invoiceRepository) GetAllInvoice(ctx context.Context, req *dto.GetAllInvoicesParam) (*[]dto.GetInvoiceResponse, error) {
	getAllInvoiceBuilder := squirrel.StatementBuilder

	if req.CustomerID != 0 {
		getAllInvoiceBuilder = getAllInvoiceBuilder.Where(squirrel.Eq{"i.customer_id": req.CustomerID})
	} else if req.MerchantID != 0 {
		getAllInvoiceBuilder = getAllInvoiceBuilder.Where(squirrel.Eq{"i.merchant_id": req.MerchantID})
	}

	getAllInvoiceSQL, args, err := getAllInvoiceBuilder.
		Select(
			"i.id AS invoice_id", "i.customer_id AS customer_id", "i.merchant_id AS merchant_id",
			"i.payment_status_id AS payment_status_id", "i.payment_type_id AS payment_type_id", "i.merchant_bank_id AS merchant_bank_id",
			"i.due_at AS due_at", "i.created_at AS created_at", "i.updated_at AS updated_at",
			"SUM(id.price * id.quantity) AS total_price", "SUM(id.quantity) AS product_quantity",
		).
		From("invoices AS i").
		InnerJoin("invoice_details AS id ON id.invoice_id = i.id").
		GroupBy("id.invoice_id").
		OrderBy("created_at DESC").
		ToSql()
	if err != nil {
		r.log.Warningln("[GetAllInvoice] Failed on build get all invoice SQL:", err.Error())
		return nil, err
	}

	getInvoiceDetails, _, err := squirrel.
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

	prepInvoiceDetail, err := r.db.PreparexContext(ctx, getInvoiceDetails)
	if err != nil {
		r.log.Warningln("[GetAllInvoice] Failed on prepared detail invoice statement:", err.Error())
		return nil, err
	}
	defer prepInvoiceDetail.Close()

	invoices := []dto.GetInvoiceResponse{}

	err = r.db.SelectContext(ctx, &invoices, getAllInvoiceSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return &invoices, nil
		}
		r.log.Warningln("[GetAllInvoice] Failed on executing get all invoice sql:", err.Error())
		return nil, err
	}

	for invoiceIndex := range invoices {
		invoices[invoiceIndex].InvoiceDetail = []dto.GetDetailInvoice{}

		err = r.db.SelectContext(ctx, &invoices[invoiceIndex].InvoiceDetail, getInvoiceDetails, invoices[invoiceIndex].InvoiceID)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			r.log.Warningln("[GetAllInvoice] Failed on executing get detail invoice sql:", err.Error())
			return nil, err
		}
	}

	return &invoices, nil
}
