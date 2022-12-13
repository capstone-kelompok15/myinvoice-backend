package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *invoiceRepository) GetInvoiceByID(ctx context.Context, invoiceID int) (*dto.GetInvoiceByID, error) {
	SQL, args, err := squirrel.
		Select("customer_id", "merchant_id").
		From("invoices").
		Where(squirrel.Eq{"id": invoiceID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetInvoiceByID] Failed on build sql:", err.Error())
		return nil, err
	}

	var data dto.GetInvoiceByID
	err = r.db.Get(&data, SQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[GetInvoiceByID] Failed on executing query:", err.Error())
		return nil, err
	}

	return &data, nil
}
