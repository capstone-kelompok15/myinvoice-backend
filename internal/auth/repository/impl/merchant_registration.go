package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *customerRepository) MerchantRegistration(ctx context.Context, req *dto.MerchantRegisterRequest) error {
	// Begin the transaction
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		r.log.Warningln("[AdminRegistration] Error while begin transaction:", err.Error())
		return err
	}

	defer func() error {
		err := tx.Rollback()
		if err != nil {
			r.log.Warningln("[AdminRegistration] Error while rollback the transaction:", err.Error())
		}
		return err
	}()

	// Insert Merchant
	insertMerchantSQL, arg1, err := squirrel.
		Insert("merchants").
		Columns("name").
		Values(req.MerchantName).
		ToSql()
	if err != nil {
		r.log.Warningln("[AdminRegistration] Error while create merchant sql from squirrel:", err.Error())
		return err
	}

	res, err := tx.ExecContext(ctx, insertMerchantSQL, arg1...)
	if err != nil {
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
		Columns("merchant_id", "address", "phone_number").
		Values(merchantID, req.MerchantAddress, req.MerchantPhoneNumber).
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
		Columns("merchant_id", "username", "password", "email", "is_verified").
		Values(merchantID, req.Username, req.Password, req.Email, false).
		ToSql()
	if err != nil {
		r.log.Warningln("[AdminRegistration] Error while create merchant admin sql from squirrel:", err.Error())
		return err
	}

	_, err = tx.ExecContext(ctx, insertMerchantAdmin, arg4...)
	if err != nil {
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