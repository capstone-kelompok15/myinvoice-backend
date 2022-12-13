package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *merchantRepository) GetMerchantProfile(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
	getMerchantBankSQL, args, err := squirrel.
		Select(
			"a.id as id",
			"a.username as username",
			"a.email as email",
			"m.merchant_name as merchant_name",
			"md.merchant_address as merchant_address",
			"md.display_profile_url as display_profile_url",
			"md.merchant_phone_number as merchant_phone_number",
		).
		From("admins AS a").
		InnerJoin("merchants AS m ON a.merchant_id = m.id").
		InnerJoin("merchant_details AS md ON a.id = md.merchant_id").
		Where(squirrel.Eq{"a.id": merchantID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetMerchantProfile] Error on creating squirrel builder", err.Error())
		return nil, err
	}

	merchant := dto.MerchantProfileResponse{}
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
