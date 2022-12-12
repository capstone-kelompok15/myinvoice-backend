package router

import (
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/apiversioning"
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/customer/handler"
	"github.com/capstone-kelompok15/myinvoice-backend/internal/customer/service"
	custommiddleware "github.com/capstone-kelompok15/myinvoice-backend/pkg/middleware/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type RouterParams struct {
	E          *echo.Echo
	Log        *logrus.Entry
	Validator  *validatorutils.Validator
	Service    service.CustomerService
	Middleware custommiddleware.Middleware
}

func InitCustomerRouter(params *RouterParams) {
	customerHandler := handler.NewCustomerHandler(&handler.CustomerHandlerParams{
		Service:   params.Service,
		Log:       params.Log,
		Validator: params.Validator,
	})

	customerV1Group := params.E.Group(apiversioning.APIVersionOne + "/customers")
	customerV1Group.GET("/me", customerHandler.GetCustomerDetails(), params.Middleware.CustomerMustAuthorized())
	customerV1Group.GET("", customerHandler.GetAllCustomer(), params.Middleware.AdminMustAuthorized())
	customerV1Group.GET("/notifications", customerHandler.GetAllNotificationCustomer(), params.Middleware.CustomerMustAuthorized())
	customerV1Group.PATCH("/me/picture", customerHandler.UpdateCustomerProfilePicture(), params.Middleware.CustomerMustAuthorized())
	customerV1Group.PUT("/me", customerHandler.UpdateCustomer(), params.Middleware.CustomerMustAuthorized())
	customerV1Group.PUT("/notifications/:id", customerHandler.MarkNotifCustomerAsRead(), params.Middleware.CustomerMustAuthorized())
}
