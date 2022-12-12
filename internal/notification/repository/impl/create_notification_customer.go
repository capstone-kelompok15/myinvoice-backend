package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *notificationRepository) CreateNotificationCustomer(ctx context.Context, req *dto.CreateNotification) error {
	insertNotifSQL, args, err := squirrel.
		Insert("customer_notifications").
		Columns("to_customer_id", "from_merchant_id", "invoice_id", "notification_title_id").
		Values(req.CustomerID, req.MerchantID, req.InvoiceID, req.NotificationTitleID).
		ToSql()
	if err != nil {
		r.log.Warningln("[CreateNotificationCustomer] Error while creating sql from squirrel", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, insertNotifSQL, args...)
	if err != nil {

		r.log.Warningln("[CreateNotificationCustomer] Error while exec the query", err.Error())
		return err
	}

	return nil
}
