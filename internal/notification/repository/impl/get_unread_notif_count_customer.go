package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *notificationRepository) GetUnreadNotifCountCustomer(ctx context.Context, customerID int) (int, error) {
	SQL, args, err := squirrel.
		Select("COUNT(id)").
		From("customer_notifications").
		Where(squirrel.Eq{"is_read": false}).
		Where(squirrel.Eq{"to_customer_id": customerID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetUnreadNotifCountMerchant] Error while creating sql from squirrel", err.Error())
		return 0, err
	}

	var count int
	err = r.db.Get(&count, SQL, args...)
	if err != nil {
		r.log.Warningln("[GetUnreadNotifCountMerchant] Error while exec the query", err.Error())
		return 0, err
	}

	return count, nil
}
