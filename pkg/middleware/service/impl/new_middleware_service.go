package impl

import (
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/middleware/repository"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type middleware struct {
	config         *config.Config
	redis          *redis.Client
	log            *logrus.Entry
	middlewareRepo repository.MiddlewareRepository
}

type MiddlewareParams struct {
	Config         *config.Config
	Redis          *redis.Client
	Log            *logrus.Entry
	MiddlewareRepo repository.MiddlewareRepository
}

func NewServiceMiddleware(params *MiddlewareParams) *middleware {
	return &middleware{
		config:         params.Config,
		redis:          params.Redis,
		log:            params.Log,
		middlewareRepo: params.MiddlewareRepo,
	}
}
