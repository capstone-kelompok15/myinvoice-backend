package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *invoiceRepository) UpdateMerchantBankID(ctx context.Context, invoiceID int, merchantBankID int) error {

	SQL, args, err := squirrel.
		Update("invoices").
		Set("merchant_bank_id", merchantBankID).
		Where(squirrel.Eq{"id": invoiceID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[UpdateMerchantBankID] Error on build sql on update", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, SQL, args...)
	if err != nil {
		r.log.Warningln("[UpdateMerchantBankID] Error on exec:", err.Error())
		return err
	}

	return nil
}
