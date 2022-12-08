package impl

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type customerRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

type CustomerRepositoryParams struct {
	DB  *sqlx.DB
	Log *logrus.Entry
}

func NewCustomerRepository(params *CustomerRepositoryParams) *customerRepository {
	return &customerRepository{
		db:  params.DB,
		log: params.Log,
	}
}
