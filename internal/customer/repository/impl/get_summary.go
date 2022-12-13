package impl

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/dateutils"
)

func (r *customerRepository) GetSummary(ctx context.Context, customerID int) (*dto.CustomerSummary, error) {
	now := time.Now().AddDate(0, -1, 0)
	oneMonthAgo := dateutils.TimeToDateStr(now)

	getSummarySQL, _, err := squirrel.
		Select("SUM(id.price * id.quantity)").
		From("invoices AS i").
		InnerJoin("invoice_details AS id ON id.invoice_id = i.id").
		Where(squirrel.GtOrEq{"i.created_at": oneMonthAgo}).
		Where(squirrel.Eq{"i.payment_status_id": "?"}).
		Where(squirrel.Eq{"i.customer_id": "?"}).
		GroupBy("i.customer_id").
		ToSql()
	if err != nil {
		r.log.Warningln("[GetSummary] Error while creating sql from squirrel", err.Error())
		return nil, err
	}

	errChan := make(chan error)
	summary := dto.CustomerSummary{}

	// Unpaid
	go func() {
		err := r.db.GetContext(ctx, &summary.TotalUnpaid, getSummarySQL, oneMonthAgo, 1, customerID)
		if err == sql.ErrNoRows {
			err = nil
			summary.TotalUnpaid = 0
		}
		errChan <- err
	}()

	// Paid
	go func() {
		err := r.db.GetContext(ctx, &summary.TotalPaid, getSummarySQL, oneMonthAgo, 3, customerID)
		if err == sql.ErrNoRows {
			err = nil
			summary.TotalPaid = 0
		}
		errChan <- err
	}()

	for i := 0; i < 2; i++ {
		err := <-errChan
		if err != nil {
			r.log.Warningln("[GetSummary] Error on executing query:", err.Error())
			return nil, err
		}
	}

	return &summary, nil
}
