package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *merchantRepository) UpdateMerchantBank(ctx context.Context, req *dto.UpdateMerchantBankDataRequest) error {
	updateBankMerchantSQL, args, err := squirrel.
		Update("merchant_banks").
		SetMap(map[string]interface{}{
			"bank_id":      req.BankID,
			"on_behalf_of": req.OnBehalfOf,
			"bank_number":  req.BankNumber,
		}).
		Where(squirrel.Eq{"id": req.MerchantBankID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[UpdateMerchantBank] Failed on build sql:", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, updateBankMerchantSQL, args...)
	if err != nil {
		r.log.Warningln("[UpdateMerchantBank] Failed on execute sql:", err.Error())
		return err
	}

	return nil
}
