package impl

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (r *merchantRepository) UpdateProfilePicture(ctx context.Context, merchantID *int, newProfilePictureURL *string) error {
	updateProfilePicutreURLSQL, args, err := squirrel.
		Update("merchant_details").
		Set("display_profile_url", *newProfilePictureURL).
		Where(squirrel.Eq{"merchant_id": *merchantID}).
		ToSql()
	if err != nil {
		if err == sql.ErrNoRows {
			return customerrors.ErrRecordNotFound
		}
		r.log.Warningln("[UpdateProfilePicture] Error while creating sql from squirrel", err.Error())
		return err
	}

	_, err = r.db.ExecContext(ctx, updateProfilePicutreURLSQL, args...)
	if err != nil {
		r.log.Warningln("[UpdateProfilePicture] Error while executing query", err.Error())
		return err
	}

	return nil
}
