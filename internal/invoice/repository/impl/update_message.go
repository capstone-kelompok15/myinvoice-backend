package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *invoiceRepository) UpdateMessage(ctx context.Context, invoiceID int, message string) error {

	SQL, args, err := squirrel.
		Update("invoices").
		Set("message", message).
		Where(squirrel.Eq{"id": invoiceID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[UpdateMessage] Error on build sql on update", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, SQL, args...)
	if err != nil {
		r.log.Warningln("[UpdateMessage] Error on exec:", err.Error())
		return err
	}

	return nil
}
