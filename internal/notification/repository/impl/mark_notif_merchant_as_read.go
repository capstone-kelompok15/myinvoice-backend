package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *notificationRepository) MarkNotifMerchantAsRead(ctx context.Context, NotifID int, MerchantID int) error {
	updateSQL, args, err := squirrel.
		Update("merchant_notifications").
		Set("is_read", true).
		Where(squirrel.Eq{"id": NotifID}).
		Where(squirrel.Eq{"to_merchant_id": MerchantID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[MarkNotifMerchantAsRead] Error while creating sql from squirrel", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, updateSQL, args...)
	if err != nil {
		r.log.Warningln("[MarkNotifMerchantAsRead] Error while executing query", err.Error())
		return err
	}

	return nil
}
