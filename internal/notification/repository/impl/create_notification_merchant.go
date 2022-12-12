package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *notificationRepository) CreateNotificationMerchant(ctx context.Context, req *dto.CreateNotification) error {
	insertNotifSQL, args, err := squirrel.
		Insert("merchant_notifications").
		Columns("to_merchant_id", "from_customer_id", "invoice_id", "notification_title_id").
		Values(req.MerchantID, req.CustomerID, req.InvoiceID, req.NotificationTitleID).
		ToSql()
	if err != nil {
		r.log.Warningln("[CreateNotificationMerchant] Error while creating sql from squirrel", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, insertNotifSQL, args...)
	if err != nil {

		r.log.Warningln("[CreateNotificationMerchant] Error while exec the query", err.Error())
		return err
	}

	return nil
}
