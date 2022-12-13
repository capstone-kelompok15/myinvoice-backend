package impl

import (
	"context"
	"database/sql"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *invoiceRepository) ValidateInvoiceID(ctx context.Context, customerID int, invoiceID int) error {
	var id int
	err := r.db.GetContext(ctx, &id, `
			SELECT id
			FROM invoices
			WHERE id = ? AND customer_id = ?
		`, invoiceID, customerID)

	if err != nil {
		if err == sql.ErrNoRows {
			return customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[ValidateInvoiceID] Error on executing query:", err.Error())
		return err
	}
	return nil
}
