package impl

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *notificationRepository) GetTitleID(title string) (int, error) {
	SQL, args, err := squirrel.
		Select("id").
		From("notification_titles").
		Where(squirrel.Eq{"name": title}).
		ToSql()
	if err != nil {
		r.log.Warningln("[GetTitleID] Error while creating sql from squirrel", err.Error())
		return 0, err
	}

	var id int
	err = r.db.Get(&id, SQL, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[GetTitleID] Error while exec the query", err.Error())
		return 0, err
	}

	return id, nil
}
