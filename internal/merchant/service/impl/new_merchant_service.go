package impl

import (
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	"github.com/capstone-kelompok15/myinvoice-backend/internal/merchant/repository"
	repoNotif "github.com/capstone-kelompok15/myinvoice-backend/internal/notification/repository"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/sirupsen/logrus"
)

type merchantService struct {
	repoNotif  repoNotif.NotificationRepository
	repo       repository.MerchantRepository
	log        *logrus.Entry
	config     *config.Config
	cloudinary *cloudinary.Cloudinary
}

type MerchantServiceParams struct {
	RepoNotif  repoNotif.NotificationRepository
	Repo       repository.MerchantRepository
	Log        *logrus.Entry
	Config     *config.Config
	Cloudinary *cloudinary.Cloudinary
}

func NewMerchantService(params *MerchantServiceParams) *merchantService {
	return &merchantService{
		repoNotif:  params.RepoNotif,
		repo:       params.Repo,
		log:        params.Log,
		config:     params.Config,
		cloudinary: params.Cloudinary,
	}
}
