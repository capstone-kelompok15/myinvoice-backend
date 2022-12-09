package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *merchantRepository) GetDashboardInvoiceOverview(ctx context.Context, merchantID int) (*dto.OverviewMerchantDashboard, error) {
	invoiceOverviewBaseBuilder := squirrel.StatementBuilder.Where(squirrel.Eq{"merchant_id": merchantID})

	getInvoiceQuantitySQL, arg1, err := invoiceOverviewBaseBuilder.
		Select("COUNT(*)").
		From("invoices").
		ToSql()
	if err != nil {
		r.log.Warningln("[GetDashboardInvoiceOverview] Failed on creating squirrel builder on get invoice quantity:", err.Error())
		return nil, err
	}

	getCustomerQuantitySQL, arg2, err := invoiceOverviewBaseBuilder.
		Select("COUNT(customer_id)").
		From("invoices").
		GroupBy("customer_id").
		ToSql()
	if err != nil {
		r.log.Warningln("[GetDashboardInvoiceOverview] Failed on creating squirrel builder on get customer quantity:", err.Error())
		return nil, err
	}

	getUnpaidInvoicesSQL, arg3, err := invoiceOverviewBaseBuilder.
		Select("COUNT(*)").
		From("invoices").
		Where(squirrel.Eq{"payment_status_id": 1}).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetDashboardInvoiceOverview] Failed on creating squirrel builder on get unpaid invoice quantity:", err.Error())
		return nil, err
	}

	getPaymentReceivedSQL, arg4, err := invoiceOverviewBaseBuilder.
		Select("SUM(id.price)").
		From("invoices as i").
		InnerJoin("invoice_details as id on id.invoice_id = i.id").
		GroupBy("id.invoice_id").
		Where(squirrel.Eq{"i.payment_status_id": 1}).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetDashboardInvoiceOverview] Failed on creating squirrel builder on get payment received quantity:", err.Error())
		return nil, err
	}

	var overview dto.OverviewMerchantDashboard
	errChan := make(chan error, 4)

	// Get Invoice Quantity SQL
	go func() {
		var invoiceQuantity *int
		err := r.db.GetContext(ctx, &invoiceQuantity, getInvoiceQuantitySQL, arg1...)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			} else {
				r.log.Warningln("[GetDashboardInvoiceOverview] Error on execute invoice quantity:", err.Error())
			}
		}

		if invoiceQuantity == nil {
			overview.InvoiceQuantity = 0
		} else {
			overview.InvoiceQuantity = *invoiceQuantity
		}
		errChan <- err
	}()

	// Get Customer Quantity
	go func() {
		var customerQuantity *int
		err := r.db.GetContext(ctx, &customerQuantity, getCustomerQuantitySQL, arg2...)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			} else {
				r.log.Warningln("[GetDashboardInvoiceOverview] Error on execute customer quantity:", err.Error())
			}
		}

		if customerQuantity == nil {
			overview.CustomerQuantity = 0
		} else {
			overview.CustomerQuantity = *customerQuantity
		}
		errChan <- err
	}()

	// Get Unpaid
	go func() {
		var unpaidInvoiceQuantity *int
		err := r.db.GetContext(ctx, &unpaidInvoiceQuantity, getUnpaidInvoicesSQL, arg3...)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			} else {
				r.log.Warningln("[GetDashboardInvoiceOverview] Error on execute unpaid invoice quantity:", err.Error())
			}
		}

		if unpaidInvoiceQuantity == nil {
			overview.UnpaidInvoiceQuantity = 0
		} else {
			overview.UnpaidInvoiceQuantity = *unpaidInvoiceQuantity
		}
		errChan <- err
	}()

	// Get Payment Received
	go func() {
		var paymentReceivedQuantity *int64
		err := r.db.GetContext(ctx, &paymentReceivedQuantity, getPaymentReceivedSQL, arg4...)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			} else {
				r.log.Warningln("[GetDashboardInvoiceOverview] Error on execute payment received quantity:", err.Error())
			}
		}

		if paymentReceivedQuantity == nil {
			overview.PaymentReceivedQuantity = 0
		} else {
			overview.PaymentReceivedQuantity = *paymentReceivedQuantity
		}
		errChan <- err
	}()

	for i := 0; i < 4; i++ {
		err := <-errChan
		if err != nil {
			return nil, err
		}
		continue
	}

	return &overview, nil
}
