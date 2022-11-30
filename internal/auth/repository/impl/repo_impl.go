package impl

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type customerRepository struct {
	db                  *sqlx.DB
	log                 *logrus.Entry
	squirrelBaseBuilder squirrel.StatementBuilderType
}

type CustomerRepositoryParams struct {
	DB  *sqlx.DB
	Log *logrus.Entry
}

func NewCustomerRepository(params *CustomerRepositoryParams) *customerRepository {
	selectActiveCustomer := squirrel.StatementBuilder.Where(squirrel.Eq{"is_inactive": 0})
	return &customerRepository{
		db:                  params.DB,
		squirrelBaseBuilder: selectActiveCustomer,
		log:                 params.Log,
	}
}
