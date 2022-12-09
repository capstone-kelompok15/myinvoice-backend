package impl

import (
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	"github.com/capstone-kelompok15/myinvoice-backend/internal/merchant/repository"
	"github.com/sirupsen/logrus"
)

type merchantService struct {
	repo   repository.MerchantRepository
	log    *logrus.Entry
	config *config.Config
}

type MerchantServiceParams struct {
	Repo   repository.MerchantRepository
	Log    *logrus.Entry
	Config *config.Config
}

func NewMerchantService(params *MerchantServiceParams) *merchantService {
	return &merchantService{
		repo:   params.Repo,
		log:    params.Log,
		config: params.Config,
	}
}
