package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *notificationRepository) MarkNotifCustomerAsRead(ctx context.Context, NotifID int, CustomerID int) error {
	updateSQL, args, err := squirrel.
		Update("customer_notifications").
		Set("is_read", true).
		Where(squirrel.Eq{"id": NotifID}).
		Where(squirrel.Eq{"to_customer_id": CustomerID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[MarkNotifCustomerAsRead] Error while creating sql from squirrel", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, updateSQL, args...)

	if err != nil {
		r.log.Warningln("[MarkNotifCustomerAsRead] Error while executing query", err.Error())
		return err
	}

	return nil
}
