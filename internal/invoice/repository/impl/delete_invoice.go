package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *invoiceRepository) DeleteInvoice(ctx context.Context, req *dto.DeleteInvoice) error {
	deleteAllInvoiceDetailsSQL, args, err := squirrel.
		Delete("invoice_details").
		Where(squirrel.Eq{"invoice_id": req.InvoiceID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[DeleteInvoice] Error on build delete invoice details sql:", err.Error())
		return err
	}

	deleteAllInvoiceSQL, args1, err := squirrel.
		Delete("invoices").
		Where(squirrel.Eq{"id": req.InvoiceID}).
		Where(squirrel.Eq{"customer_id": req.CustomerID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[DeleteInvoice] Error on build delete invoices sql:", err.Error())
		return err
	}

	deleteAllCustomerNotificationsSQL, args2, err := squirrel.
		Delete("customer_notifications").
		Where(squirrel.Eq{"invoice_id": req.InvoiceID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[DeleteInvoice] Error on build delete customer notifications sql:", err.Error())
		return err
	}

	deleteAllMerchantNotificationsSQL, args3, err := squirrel.
		Delete("merchant_notifications").
		Where(squirrel.Eq{"invoice_id": req.InvoiceID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[DeleteInvoice] Error on build delete merchant notifications sql:", err.Error())
		return err
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		r.log.Warningln("[DeleteInvoice] Error on create database transaction:", err.Error())
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, deleteAllInvoiceDetailsSQL, args...)
	if err != nil {
		r.log.Warningln("[DeleteInvoice] Error on deleting invoice details:", err.Error())
		return err
	}

	_, err = tx.ExecContext(ctx, deleteAllCustomerNotificationsSQL, args2...)
	if err != nil {
		r.log.Warningln("[DeleteInvoice] Error on deleting customer notifications:", err.Error())
		return err
	}

	_, err = tx.ExecContext(ctx, deleteAllMerchantNotificationsSQL, args3...)
	if err != nil {
		r.log.Warningln("[DeleteInvoice] Error on deleting merchant notifications:", err.Error())
		return err
	}

	_, err = tx.ExecContext(ctx, deleteAllInvoiceSQL, args1...)
	if err != nil {
		r.log.Warningln("[DeleteInvoice] Error on deleting invoices:", err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		r.log.Warningln("[DeleteInvoice] Error on commiting to the database:", err.Error())
		return err
	}

	return nil
}
