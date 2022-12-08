package impl

import (
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	"github.com/capstone-kelompok15/myinvoice-backend/internal/auth/repository"
	"github.com/go-redis/redis/v8"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/sirupsen/logrus"
)

type authService struct {
	repo    repository.AuthRepository
	log     *logrus.Entry
	redis   *redis.Client
	mailgun *mailgun.MailgunImpl
	config  *config.Config
}

type AuthServiceParams struct {
	Repo    repository.AuthRepository
	Log     *logrus.Entry
	Redis   *redis.Client
	Mailgun *mailgun.MailgunImpl
	Config  *config.Config
}

func NewAuthService(params *AuthServiceParams) *authService {
	return &authService{
		repo:    params.Repo,
		log:     params.Log,
		redis:   params.Redis,
		mailgun: params.Mailgun,
		config:  params.Config,
	}
}
