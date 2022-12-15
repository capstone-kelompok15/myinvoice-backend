package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *invoiceRepository) GetMerchantProfile(ctx context.Context, invoiceID int) (*dto.MerchantBriefDate, error) {
	getMerchantBankSQL, args, err := squirrel.
		Select(
			"a.username as username",
			"a.email as email",
			"m.merchant_name as merchant_name",
			"m.id as merchant_id",
		).
		From("admins AS a").
		InnerJoin("merchants AS m ON a.merchant_id = m.id").
		InnerJoin("invoices AS i ON i.merchant_id = m.id").
		Where(squirrel.Eq{"i.id": invoiceID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetMerchantProfile] Error on creating squirrel builder", err.Error())
		return nil, err
	}

	merchant := dto.MerchantBriefDate{}
	err = r.db.Get(&merchant, getMerchantBankSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[GetMerchantProfile] Error on creating squirrel builder", err.Error())
		return nil, err
	}

	return &merchant, nil
}
