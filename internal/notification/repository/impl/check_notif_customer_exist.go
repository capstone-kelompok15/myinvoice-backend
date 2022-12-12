package impl

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *notificationRepository) CheckNotifCustomerExist(ID int) error {
	SQL, args, err := squirrel.
		Select("id").
		From("customer_notifications").
		Where(squirrel.Eq{"id": ID}).
		ToSql()
	if err != nil {
		r.log.Warningln("[CheckNotifCustomerExist] Error while creating sql from squirrel", err.Error())
		return err
	}
	var id int
	err = r.db.Get(&id, SQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[CheckNotifCustomerExist] Error while exec the query", err.Error())
		return err
	}

	return nil
}
