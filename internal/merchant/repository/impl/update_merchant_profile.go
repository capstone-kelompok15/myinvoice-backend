package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *merchantRepository) UpdateMerchantProfile(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
	updateUsernameSQL, args1, err := squirrel.
		Update("admins").
		Set("username", req.Username).
		Where(squirrel.Eq{"id": *merchantID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[UpdateMerchantProfile] Error while creating sql from squirrel", err.Error())
		return err
	}
	updateMerchantSQL, args2, err := squirrel.
		Update("merchants").
		Set("merchant_name", req.MerchantName).
		Where(squirrel.Eq{"id": *merchantID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[UpdateMerchantProfile] Error while creating sql from squirrel", err.Error())
		return err
	}
	updateMerchantDetailsSQL, args3, err := squirrel.
		Update("merchant_details").
		Set("merchant_address", req.MerchantAddress).
		Set("merchant_phone_number", req.MerchantPhoneNumber).
		Where(squirrel.Eq{"merchant_id": *merchantID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[UpdateMerchantProfile] Error while creating sql from squirrel", err.Error())
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.log.Warningln("[UpdateMerchantProfile] Error while execute ", err.Error())
		return err
	}

	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, updateUsernameSQL, args1...)
	if err != nil {
		r.log.Warningln("[UpdateMerchantProfile] Error while executing query", err.Error())
		return err
	}
	_, err = tx.ExecContext(ctx, updateMerchantSQL, args2...)
	if err != nil {
		r.log.Warningln("[UpdateMerchantProfile] Error while executing query", err.Error())
		return err
	}
	_, err = tx.ExecContext(ctx, updateMerchantDetailsSQL, args3...)
	if err != nil {
		r.log.Warningln("[UpdateMerchantProfile] Error while executing query", err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		r.log.Warningln("[UpdateMerchantProfile] Error while commit the transaction", err.Error())
		return err
	}

	return nil
}
