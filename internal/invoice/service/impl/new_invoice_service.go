package impl

import (
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	"github.com/capstone-kelompok15/myinvoice-backend/internal/invoice/repository"
	repoNotif "github.com/capstone-kelompok15/myinvoice-backend/internal/notification/repository"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-redis/redis/v8"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/sirupsen/logrus"
)

type invoiceService struct {
	repoNotif  repoNotif.NotificationRepository
	repo       repository.InvoiceRepository
	log        *logrus.Entry
	redis      *redis.Client
	mailgun    *mailgun.MailgunImpl
	config     *config.Config
	cloudinary *cloudinary.Cloudinary
}

type InvoiceService struct {
	RepoNotif  repoNotif.NotificationRepository
	Repo       repository.InvoiceRepository
	Log        *logrus.Entry
	Redis      *redis.Client
	Mailgun    *mailgun.MailgunImpl
	Config     *config.Config
	Cloudinary *cloudinary.Cloudinary
}

func NewInvoiceService(params *InvoiceService) *invoiceService {
	return &invoiceService{
		repoNotif:  params.RepoNotif,
		repo:       params.Repo,
		log:        params.Log,
		redis:      params.Redis,
		mailgun:    params.Mailgun,
		config:     params.Config,
		cloudinary: params.Cloudinary,
	}
}
