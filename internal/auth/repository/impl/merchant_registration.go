package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/go-sql-driver/mysql"
)

func (r *authRepository) MerchantRegistration(ctx context.Context, req *dto.MerchantRegisterRequest) error {
	// Begin the transaction
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		r.log.Warningln("[AdminRegistration] Error while begin transaction:", err.Error())
		return err
	}
	defer tx.Rollback()

	// Insert Merchant
	insertMerchantSQL, arg1, err := squirrel.
		Insert("merchants").
		Columns("merchant_name").
		Values(req.MerchantName).
		ToSql()
	if err != nil {
		r.log.Warningln("[AdminRegistration] Error while create merchant sql from squirrel:", err.Error())
		return err
	}

	res, err := tx.ExecContext(ctx, insertMerchantSQL, arg1...)
	if err != nil {
		if errCode, ok := err.(*mysql.MySQLError); ok {
			if errCode.Number == 1062 {
				return customerrors.ErrUniqueRecord
			}
		}
		r.log.Warningln("[AdminRegistration] Error while insert merchant:", err.Error())
		return err
	}

	merchantID, err := res.LastInsertId()
	if err != nil {
		r.log.Warningln("[AdminRegistration] Error while getting the last inserted ID (merchant ID):", err.Error())
		return err
	}

	// Insert merchant detail
	insertMerchantDetailSQL, arg2, err := squirrel.
		Insert("merchant_details").
		Columns("merchant_id", "merchant_address").
		Values(merchantID, req.MerchantAddress).
		ToSql()
	if err != nil {
		r.log.Warningln("[AdminRegistration] Error while create merchant detail sql from squirrel:", err.Error())
		return err
	}

	_, err = tx.ExecContext(ctx, insertMerchantDetailSQL, arg2...)
	if err != nil {
		r.log.Warningln("[AdminRegistration] Error while insert merchant details:", err.Error())
		return err
	}

	// Insert merchant banks
	insertMerchantBanksBuilder := squirrel.
		Insert("merchant_banks").
		Columns("merchant_id", "bank_id", "on_behalf_of", "bank_number")

	for _, bank := range req.MerchantBank {
		insertMerchantBanksBuilder = insertMerchantBanksBuilder.
			Values(merchantID, bank.BankID, bank.OnBehalfOf, bank.BankNumber)
	}

	insertMerchantBanksSQL, arg3, err := insertMerchantBanksBuilder.ToSql()
	if err != nil {
		r.log.Warningln("[AdminRegistration] Error while create merchant banks sql from squirrel:", err.Error())
		return err
	}

	_, err = tx.ExecContext(ctx, insertMerchantBanksSQL, arg3...)
	if err != nil {
		r.log.Warningln("[AdminRegistration] Error while insert merchant banks:", err.Error())
		return err
	}

	// Create merchant admin
	insertMerchantAdmin, arg4, err := squirrel.
		Insert("admins").
		Columns("merchant_id", "username", "admin_password", "email", "is_verified").
		Values(merchantID, req.Username, req.Password, req.Email, false).
		ToSql()
	if err != nil {
		r.log.Warningln("[AdminRegistration] Error while create merchant admin sql from squirrel:", err.Error())
		return err
	}

	_, err = tx.ExecContext(ctx, insertMerchantAdmin, arg4...)
	if err != nil {
		if errCode, ok := err.(*mysql.MySQLError); ok {
			if errCode.Number == 1062 {
				return customerrors.ErrMerchantNameDuplicated
			}
		}
		r.log.Warningln("[AdminRegistration] Error while insert merchant admin:", err.Error())
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		r.log.Warningln("[AdminRegistration] Error while committing the transaction:", err.Error())
		return err
	}

	return nil
}
