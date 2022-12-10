package router

import (
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/apiversioning"
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/auth/handler"
	"github.com/capstone-kelompok15/myinvoice-backend/internal/auth/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type RouterParams struct {
	E         *echo.Echo
	Log       *logrus.Entry
	Validator *validatorutils.Validator
	Service   service.AuthService
}

func InitAuthRouter(params *RouterParams) {
	authHandler := handler.NewAuthHandler(&handler.AuthHandlerParams{
		Service:   params.Service,
		Log:       params.Log,
		Validator: params.Validator,
	})

	authV1Group := params.E.Group(apiversioning.APIVersionOne + "/auth")
	authV1Group.POST("/register/customer", authHandler.RegisterCustomer())
	authV1Group.POST("/register/merchant", authHandler.RegisterMerchant())
	authV1Group.POST("/verification/customer", authHandler.CustomerEmailVerification())
	authV1Group.POST("/verification/admin", authHandler.AdminEmailVerification())
	authV1Group.POST("/verification/admin/refresh", authHandler.RefreshAdminEmailVerificationCode())
	authV1Group.POST("/verification/customer/refresh", authHandler.RefreshEmailVerificationCode())
	authV1Group.POST("/login/customer", authHandler.CustomerLogin())
	authV1Group.POST("/login/admin", authHandler.AdminLogin())
	authV1Group.POST("/admin/refresh", authHandler.AdminRefreshToken())
	authV1Group.POST("/reset/password/customer/request", authHandler.CustomerResetPasswordRequest())
	authV1Group.POST("/reset/password/customer", authHandler.CustomerResetPassword())
	authV1Group.POST("/reset/password/admin/request", authHandler.AdminResetPasswordRequest())
}
