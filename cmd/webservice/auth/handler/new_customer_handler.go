package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/internal/auth/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/websocketutils"
	"github.com/sirupsen/logrus"
)

type authHandler struct {
	service       service.AuthService
	log           *logrus.Entry
	validator     *validatorutils.Validator
	websocketPool *websocketutils.Pool
}

type AuthHandlerParams struct {
	Service       service.AuthService
	Log           *logrus.Entry
	Validator     *validatorutils.Validator
	WebsocketPool *websocketutils.Pool
}

func NewAuthHandler(params *AuthHandlerParams) *authHandler {
	return &authHandler{
		service:       params.Service,
		log:           params.Log,
		validator:     params.Validator,
		websocketPool: params.WebsocketPool,
	}
}
