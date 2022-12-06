package impl

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type bankRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

type BankRepositoryParams struct {
	DB  *sqlx.DB
	Log *logrus.Entry
}

func NewBankRepository(params *BankRepositoryParams) *bankRepository {
	return &bankRepository{
		db:  params.DB,
		log: params.Log,
	}
}
