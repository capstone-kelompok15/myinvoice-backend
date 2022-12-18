package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *invoiceRepository) ValidateInvoiceID(ctx context.Context, customerID int, invoiceID int, invoiceDetailID *int) error {
	validateInvoiceID := squirrel.
		Select("i.id").
		From("invoices AS i").
		Where(squirrel.Eq{"i.id": invoiceID}).
		Where(squirrel.Eq{"i.customer_id": customerID})

	if invoiceDetailID != nil {
		validateInvoiceID = validateInvoiceID.InnerJoin("invoice_details AS id ON id.invoice_id = i.id").
			Where(squirrel.Eq{"id.id": *invoiceDetailID})
	}

	validateInvoiceIDSQL, args, err := validateInvoiceID.ToSql()
	if err != nil {
		return err
	}

	var id int

	err = r.db.GetContext(ctx, &id, validateInvoiceIDSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[ValidateInvoiceID] Error on executing query:", err.Error())
		return err
	}
	return nil
}
