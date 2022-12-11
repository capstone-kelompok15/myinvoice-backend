package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *merchantRepository) ValidateMerchantBank(ctx context.Context, merchantID int, merchantBankID int) error {
	validateMerchantBankSQL, args, err := squirrel.
		Select("id").
		From("merchant_banks AS mb").
		Where(squirrel.Eq{"merchant_id": merchantID}).
		Where(squirrel.Eq{"id": merchantBankID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[ValidateMerchantBank] Failed on build sql:", err.Error())
		return err
	}

	var id int
	err = r.db.GetContext(ctx, &id, validateMerchantBankSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[ValidateMerchantBank] Failed execute query:", err.Error())
		return err
	}

	return nil
}
