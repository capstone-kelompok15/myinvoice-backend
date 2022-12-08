package impl

import (
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	"github.com/capstone-kelompok15/myinvoice-backend/internal/customer/repository"
	"github.com/go-redis/redis/v8"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/sirupsen/logrus"
)

type customerService struct {
	repo    repository.CustomerRepository
	log     *logrus.Entry
	redis   *redis.Client
	mailgun *mailgun.MailgunImpl
	config  *config.Config
}

type CustomerServiceParams struct {
	Repo    repository.CustomerRepository
	Log     *logrus.Entry
	Redis   *redis.Client
	Mailgun *mailgun.MailgunImpl
	Config  *config.Config
}

func NewCustomerService(params *CustomerServiceParams) *customerService {
	return &customerService{
		repo:    params.Repo,
		log:     params.Log,
		redis:   params.Redis,
		mailgun: params.Mailgun,
		config:  params.Config,
	}
}
