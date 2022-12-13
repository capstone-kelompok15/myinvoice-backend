package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *invoiceRepository) GetReport(ctx context.Context, params *dto.ReportParams) (*dto.ReportResponse, error) {
	getSummarySQL, _, err := squirrel.
		Select("SUM(id.quantity)").
		From("invoices AS i").
		InnerJoin("invoice_details AS id ON id.invoice_id = i.id").
		Where(squirrel.LtOrEq{"i.created_at": "?"}).
		Where(squirrel.GtOrEq{"i.created_at": "?"}).
		Where(squirrel.Eq{"i.payment_status_id": "?"}).
		Where(squirrel.Eq{"i.customer_id": "?"}).
		GroupBy("i.customer_id").
		ToSql()
	if err != nil {
		r.log.Warningln("[GetReport] Failed on build get summary sql:", err.Error())
		return nil, err
	}

	getOverviewTransactionSQL, _, err := squirrel.
		Select("SUM(id.price * id.quantity) AS transaction_total, SUM(id.quantity) AS transaction_quantity").
		From("invoices AS i").
		InnerJoin("invoice_details AS id ON id.invoice_id = i.id").
		Where(squirrel.LtOrEq{"i.created_at": "?"}).
		Where(squirrel.GtOrEq{"i.created_at": "?"}).
		Where(squirrel.Eq{"i.payment_status_id": "?"}).
		Where(squirrel.Eq{"i.customer_id": "?"}).
		GroupBy("i.customer_id").
		ToSql()
	if err != nil {
		r.log.Warningln("[GetReport] Failed on build get overview transaction sql:", err.Error())
		return nil, err
	}

	prep, err := r.db.PreparexContext(ctx, getSummarySQL)
	if err != nil {
		return nil, err
	}
	defer prep.Close()

	errChan := make(chan error)
	report := dto.ReportResponse{}
	report.Reports = make([]dto.ReportDate, len(params.ReportDaysInt))

	go func() {
		errChanInside := make(chan error)

		for i, date := range params.ReportDaysStr {
			go func(index int, date dto.ReportDateStr) {
				err := prep.GetContext(
					ctx, &report.Reports[index].Value,
					date.EndDate, date.StartDate, params.PaymentStatus, params.CustomerID,
				)

				report.Reports[index].Date = date.EndDate
				if err == sql.ErrNoRows {
					report.Reports[index].Value = 0
					err = nil
				}

				errChanInside <- err
			}(i, date)
		}

		for i := 0; i < len(params.ReportDaysStr); i++ {
			err := <-errChanInside
			if err != nil {
				errChan <- err
				return
			}
		}
		errChan <- nil
	}()
	go func() {
		startDate := params.ReportDaysStr[0].StartDate
		endDate := params.ReportDaysStr[len(params.ReportDaysStr)-1].EndDate
		reportTransaction := dto.ReportTransaction{}

		err := r.db.GetContext(
			ctx, &reportTransaction, getOverviewTransactionSQL,
			endDate, startDate, params.PaymentStatus, params.CustomerID,
		)

		if err == sql.ErrNoRows {
			reportTransaction = dto.ReportTransaction{
				TransactionQuantity: 0,
				TransactionTotal:    0,
			}
			err = nil
		}

		report.ReportTransaction = reportTransaction
		errChan <- err
	}()

	for i := 0; i < 2; i++ {
		err := <-errChan
		if err != nil {
			r.log.Warningln("[GetReport] Error on execute query:", err.Error())
			return nil, err
		}
	}

	return &report, nil
}
