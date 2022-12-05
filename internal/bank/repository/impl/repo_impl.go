package impl

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type bankRepository struct {
	db                  *sqlx.DB
	log                 *logrus.Entry
	squirrelBaseBuilder squirrel.StatementBuilderType
}

type BankRepositoryParams struct {
	DB  *sqlx.DB
	Log *logrus.Entry
}

func NewBankRepository(params *BankRepositoryParams) *bankRepository {
	selectActiveBank := squirrel.StatementBuilder.Where(squirrel.Eq{"is_inactive": 0})
	return &bankRepository{
		db:                  params.DB,
		squirrelBaseBuilder: selectActiveBank,
		log:                 params.Log,
	}
}
