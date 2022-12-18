package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *invoiceRepository) DeleteDetailInvoice(ctx context.Context, req *dto.DeleteInvoice) error {
	deleteInvoiceDetailsSQL, args, err := squirrel.
		Delete("invoice_details").
		Where(squirrel.Eq{"invoice_id": req.InvoiceID}).
		Where(squirrel.Eq{"id": req.InvoiceDetailID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[DeleteDetailInvoice] Error on build delete invoice details sql:", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, deleteInvoiceDetailsSQL, args...)
	if err != nil {
		r.log.Warningln("[DeleteDetailInvoice] Error on deleting detail invoices:", err.Error())
		return err
	}

	return nil
}
