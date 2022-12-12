package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *merchantRepository) GetMerchantBank(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error) {
	getMerchantBankSQL, args, err := squirrel.
		Select(
			"mb.on_behalf_of AS on_behalf_of",
			"mb.bank_number AS bank_number",
			"b.bank_name AS bank_name",
			"b.code AS bank_code").
		From("merchant_banks AS mb").
		InnerJoin("banks AS b ON b.id = mb.bank_id").
		Where(squirrel.Eq{"mb.merchant_id": req.MerchantID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetMerchantBank] Error on creating squirrel builder", err.Error())
		return nil, err
	}

	merchantBank := []dto.GetMerchantBankResponse{}
	err = r.db.Select(&merchantBank, getMerchantBankSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[GetMerchantBank] Error on creating squirrel builder", err.Error())
		return nil, err
	}

	return &merchantBank, nil
}
