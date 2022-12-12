package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *notificationRepository) GetAllNotificationCustomer(ctx context.Context, customerID int, req *dto.NotificationRequest) (*[]dto.CustomerNotificationDB, error) {
	startIndex := (req.Page - 1) * req.Limit
	getAllCustomerSQL, args, err := squirrel.
		Select("cn.id as id", "cn.is_read as is_read", "cn.created_at as created_at", "cn.to_customer_id as customer_id", "cn.from_merchant_id as merchant_id", "cn.invoice_id as invoice_id", "nt.name as title", "nts.name as type", "m.merchant_name as merchant_name").
		From("customer_notifications as cn").
		InnerJoin("merchants as m on cn.from_merchant_id = m.id").
		InnerJoin("notification_titles as nt on cn.notification_title_id = nt.id").
		InnerJoin("notification_types as nts on nt.notification_type_id = nts.id").
		Where(squirrel.Eq{"cn.to_customer_id": customerID}).
		OrderBy("cn.created_at DESC").
		Offset(startIndex).Limit(req.Limit).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetAllNotificationCustomer] Error while creating sql from squirrel", err.Error())
		return nil, err
	}

	var notifications []dto.CustomerNotificationDB
	err = r.db.Select(&notifications, getAllCustomerSQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[GetAllNotificationCustomer] Error while exec the query", err.Error())
		return nil, err
	}

	return &notifications, nil
}
