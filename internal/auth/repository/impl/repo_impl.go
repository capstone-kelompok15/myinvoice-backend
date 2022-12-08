package impl

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type authRepository struct {
	db                  *sqlx.DB
	log                 *logrus.Entry
	squirrelBaseBuilder squirrel.StatementBuilderType
}

type AuthRepositoryParams struct {
	DB  *sqlx.DB
	Log *logrus.Entry
}

func NewAuthRepository(params *AuthRepositoryParams) *authRepository {
	selectActiveCustomer := squirrel.StatementBuilder.Where(squirrel.Eq{"is_inactive": 0})
	return &authRepository{
		db:                  params.DB,
		squirrelBaseBuilder: selectActiveCustomer,
		log:                 params.Log,
	}
}
