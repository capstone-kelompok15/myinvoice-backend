package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *invoiceRepository) CreateInvoice(ctx context.Context, merchantID int, req *dto.CreateInvoiceRequest) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		r.log.Warningln("[CreateInvoice] Error on creating database transaction:", err.Error())
		return err
	}
	defer tx.Rollback()

	createInvoiceSQL, args1, err := squirrel.
		Insert("invoices").
		Columns(
			"merchant_id", "customer_id",
			"payment_status_id", "due_at", "note",
		).
		Values(
			merchantID, req.CustomerID,
			1, req.DueAt, req.Note,
		).
		ToSql()
	if err != nil {
		r.log.Warningln("[CreateInvoice] Error on build sql on create invoice:", err.Error())
		return err
	}

	res, err := tx.ExecContext(ctx, createInvoiceSQL, args1...)
	if err != nil {
		r.log.Warningln("[CreateInvoice] Error on create invoice:", err.Error())
		return err
	}

	invoiceID, err := res.LastInsertId()
	if err != nil {
		r.log.Warningln("[CreateInvoice] Error on get last insert ID:", err.Error())
		return err
	}

	createInvoiceDetails := squirrel.
		Insert("invoice_details").
		Columns("invoice_id", "product", "quantity", "price")

	for _, item := range req.Items {
		createInvoiceDetails = createInvoiceDetails.Values(invoiceID, item.Product, item.Quantity, item.Price)
	}

	createInvoiceDetailsSQL, args2, err := createInvoiceDetails.ToSql()
	if err != nil {
		r.log.Warningln("[CreateInvoice] Error on build sql on create invoice detail:", err.Error())
		return err
	}

	_, err = tx.ExecContext(ctx, createInvoiceDetailsSQL, args2...)
	if err != nil {
		r.log.Warningln("[CreateInvoice] Error executing invoice detail:", err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		r.log.Warningln("[CreateInvoice] Error on commit the transaction:", err.Error())
		return err
	}

	return nil
}
