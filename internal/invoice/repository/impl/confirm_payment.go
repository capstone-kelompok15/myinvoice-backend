package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/paymentstatusutils"
)

func (r *invoiceRepository) ConfirmPayment(ctx context.Context, invoiceID int) error {

	SQL, args, err := squirrel.
		Update("invoices").
		Set("payment_status_id", paymentstatusutils.Pending).
		Where(squirrel.Eq{"id": invoiceID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[ConfirmPayment] Error on build sql on update", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, SQL, args...)
	if err != nil {
		r.log.Warningln("[ConfirmPayment] Error on exec:", err.Error())
		return err
	}

	return nil
}
