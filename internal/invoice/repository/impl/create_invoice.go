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
			"merchant_id", "customer_id", "payment_type_id",
			"payment_status_id", "merchant_bank_id", "due_at",
		).
		Values(
			merchantID, req.CustomerID, req.PaymentType,
			1, req.MerchantBank, req.DueAt,
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

	for _, invoiceDetail := range req.InvoiceDetails {
		createInvoiceDetails = createInvoiceDetails.Values(invoiceID, invoiceDetail.Product, invoiceDetail.Quantity, invoiceDetail.Price)
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
