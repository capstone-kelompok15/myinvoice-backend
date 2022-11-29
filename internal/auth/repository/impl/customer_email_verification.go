package impl

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *customerRepository) CustomerEmailVerification(ctx context.Context, req *dto.CustomerEmailVerification) error {
	var args []interface{}

	subQuerySQL, arg1, err := squirrel.
		Select("id").
		From("customers").
		Where(squirrel.Eq{"email": req.Email}).
		ToSql()
	if err != nil {
		r.log.Warningln("[CustomerEmailVerification] Error while creating subquery sql from squirrel", err.Error())
		return err
	}

	updateVerifiedStatusSQL, arg2, err := squirrel.
		Update("user_settings").
		Set("is_verified", true).
		Where(fmt.Sprintf("customer_id = (%s)", subQuerySQL)).
		ToSql()

	args = append(args, arg2...)
	args = append(args, arg1...)

	if err != nil {
		r.log.Warningln("[CustomerEmailVerification] Error while create update sql from squirrel", err.Error())
		return err
	}

	r.db.ExecContext(ctx, updateVerifiedStatusSQL, args...)

	return nil
}
