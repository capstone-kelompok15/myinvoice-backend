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
	Service   service.CustomerService
}

func InitCustomerRouter(params *RouterParams) {
	customerHandler := handler.NewCustomerHandler(&handler.CustomerHandler{
		Service:   params.Service,
		Log:       params.Log,
		Validator: params.Validator,
	})

	customerV1Group := params.E.Group(apiversioning.APIVersionOne + "/auth")
	customerV1Group.POST("/register/customer", customerHandler.RegisterCustomer())
	customerV1Group.POST("/register/merchant", customerHandler.RegisterMerchant())
	customerV1Group.POST("/verification/customer", customerHandler.CustomerEmailVerification())
	customerV1Group.POST("/verification/refresh", customerHandler.RefreshEmailVerificationCode())
	customerV1Group.POST("/login/customer", customerHandler.CustomerLogin())
	customerV1Group.POST("/login/admin", customerHandler.AdminLogin())
	customerV1Group.POST("/admin/refresh", customerHandler.AdminRefreshToken())
	customerV1Group.POST("/reset/password/customer", customerHandler.CustomerResetPasswordRequest())
}
