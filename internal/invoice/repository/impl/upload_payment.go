package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *invoiceRepository) UploadPayment(ctx context.Context, invoiceID int, uploadedURL string) error {
	insertPaymentURLSQL, args, err := squirrel.
		Update("invoices").
		SetMap(map[string]interface{}{
			"approval_document_url": uploadedURL,
			"payment_status_id":     2,
		}).
		Where(squirrel.Eq{"id": invoiceID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[UploadPayment] Error on build sql:", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, insertPaymentURLSQL, args...)
	if err != nil {
		r.log.Warningln("[UploadPayment] Error on executing query:", err.Error())
		return err
	}

	return nil
}
