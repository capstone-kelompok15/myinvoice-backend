package impl

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type invoiceRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

type InvoiceRepositoryParams struct {
	DB  *sqlx.DB
	Log *logrus.Entry
}

func NewInvoiceRepository(params *InvoiceRepositoryParams) *invoiceRepository {
	return &invoiceRepository{
		db:  params.DB,
		log: params.Log,
	}
}
