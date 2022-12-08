package impl

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type middlewareRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

type MiddlewareRepositoryParams struct {
	DB  *sqlx.DB
	Log *logrus.Entry
}

func NewRepositoryMiddleware(params *MiddlewareRepositoryParams) *middlewareRepository {
	return &middlewareRepository{
		db:  params.DB,
		log: params.Log,
	}
}
