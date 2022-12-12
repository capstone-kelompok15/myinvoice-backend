package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *merchantRepository) CreateMerchantBank(ctx context.Context, merchantID int, req *dto.MerchantBankData) error {
	createMerchantBankSQL, args, err := squirrel.
		Insert("merchant_banks").
		Columns("merchant_id", "bank_id", "on_behalf_of", "bank_number").
		Values(merchantID, req.BankID, req.OnBehalfOf, req.BankNumber).
		ToSql()
	if err != nil {
		r.log.Warningln("[CreateMerchantBank] Failed on build sql:", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, createMerchantBankSQL, args...)
	if err != nil {
		r.log.Warningln("[CreateMerchantBank] Failed on executing sql:", err.Error())
		return err
	}

	return nil
}
