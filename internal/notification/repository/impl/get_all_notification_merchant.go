package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *notificationRepository) GetAllNotificationMerchant(ctx context.Context, customerID int, req *dto.NotificationRequest) (*[]dto.MerchantNotificationDB, error) {
	startIndex := (req.Page - 1) * req.Limit
	SQL, args, err := squirrel.
		Select("mn.id as id", "mn.is_read as is_read", "mn.created_at as created_at", "mn.to_merchant_id as customer_id", "mn.from_customer_id as merchant_id", "mn.invoice_id as invoice_id", "nt.name as title", "nts.name as type", "cd.full_name as customer_name").
		From("merchant_notifications as mn").
		InnerJoin("customer_details as cd on mn.from_customer_id = cd.customer_id").
		InnerJoin("notification_titles as nt on mn.notification_title_id = nt.id").
		InnerJoin("notification_types as nts on nt.notification_type_id = nts.id").
		Where(squirrel.Eq{"mn.to_merchant_id": customerID}).
		OrderBy("mn.created_at DESC").
		Offset(startIndex).Limit(req.Limit).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetAllNotificationMerchant] Error while creating sql from squirrel", err.Error())
		return nil, err
	}

	var notifications []dto.MerchantNotificationDB
	err = r.db.Select(&notifications, SQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[GetAllNotificationMerchant] Error while exec the query", err.Error())
		return nil, err
	}

	return &notifications, nil
}
