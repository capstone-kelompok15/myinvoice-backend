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
}
