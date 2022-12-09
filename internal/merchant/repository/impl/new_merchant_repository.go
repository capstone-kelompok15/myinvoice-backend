package impl

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type merchantRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

type MerchantRepositoryParams struct {
	DB  *sqlx.DB
	Log *logrus.Entry
}

func NewBankRepository(params *MerchantRepositoryParams) *merchantRepository {
	return &merchantRepository{
		db:  params.DB,
		log: params.Log,
	}
}
