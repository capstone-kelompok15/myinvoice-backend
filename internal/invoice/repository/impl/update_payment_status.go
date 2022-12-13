package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *invoiceRepository) UpdatePaymentStatus(ctx context.Context, invoiceID int, paymentStatusID int) error {

	SQL, args, err := squirrel.
		Update("invoices").
		Set("payment_status_id", paymentStatusID).
		Where(squirrel.Eq{"id": invoiceID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[UpdatePaymentStatus] Error on build sql on update", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, SQL, args...)
	if err != nil {
		r.log.Warningln("[UpdatePaymentStatus] Error on exec:", err.Error())
		return err
	}

	return nil
}
