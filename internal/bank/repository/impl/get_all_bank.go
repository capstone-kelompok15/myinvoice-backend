package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *bankRepository) GetAllBank(ctx context.Context) (*[]dto.BankResponse, error) {
	getAllBankSql, _, err := squirrel.
		Select("*").
		From("banks").
		ToSql()
	if err != nil {
		r.log.Warningln("[GetAllBank] Error while creating sql from squirrel", err.Error())
		return nil, err
	}

	var bankResponse []dto.BankResponse
	err = r.db.Select(&bankResponse, getAllBankSql)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[GetAllBank] Error while exec the query", err.Error())
		return nil, err
	}
	return &bankResponse, nil
}
