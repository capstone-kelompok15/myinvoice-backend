package router

import (
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/apiversioning"
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/invoice/handler"
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	"github.com/capstone-kelompok15/myinvoice-backend/internal/invoice/service"
	custommiddleware "github.com/capstone-kelompok15/myinvoice-backend/pkg/middleware/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/websocketutils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type InvoiceRouterParams struct {
	E             *echo.Echo
	Log           *logrus.Entry
	Validator     *validatorutils.Validator
	Service       service.InvoiceService
	Middleware    custommiddleware.Middleware
	WebsocketPool *websocketutils.Pool
	Config        *config.Config
}

func InitInvoiceRouter(params *InvoiceRouterParams) {
	invoiceHandler := handler.NewInvoiceHandler(&handler.InvoiceHandlerParams{
		Service:       params.Service,
		Log:           params.Log,
		Validator:     params.Validator,
		WebsocketPool: params.WebsocketPool,
		Config:        params.Config,
	})

	invoiceV1Group := params.E.Group(apiversioning.APIVersionOne + "/invoices")
	invoiceV1Group.POST("", invoiceHandler.CreateInvoice(), params.Middleware.AdminMustAuthorized())
	invoiceV1Group.GET("/merchants", invoiceHandler.MerchantGetAllInvoices(), params.Middleware.AdminMustAuthorized())
	invoiceV1Group.GET("/customers", invoiceHandler.CustomerGetAllInvoices(), params.Middleware.CustomerMustAuthorized())
	invoiceV1Group.GET("/:invoice_id/merchants", invoiceHandler.MerchantGetDetailInvoiceByID(), params.Middleware.AdminMustAuthorized())
	invoiceV1Group.GET("/:invoice_id/customers", invoiceHandler.CustomerGetDetailInvoiceByID(), params.Middleware.CustomerMustAuthorized())
	invoiceV1Group.GET("/merchants/customers", invoiceHandler.GetCustomers(), params.Middleware.AdminMustAuthorized())
	invoiceV1Group.PATCH("/:invoice_id/payments/upload", invoiceHandler.UploadPayment(), params.Middleware.CustomerMustAuthorized())
	invoiceV1Group.PUT("/:invoice_id/confirm", invoiceHandler.ConfirmPayment(), params.Middleware.CustomerMustAuthorized())
	invoiceV1Group.PUT("/:invoice_id/accept", invoiceHandler.AcceptPayment(), params.Middleware.AdminMustAuthorized())
	invoiceV1Group.PUT("/:invoice_id/reject", invoiceHandler.RejectPayment(), params.Middleware.AdminMustAuthorized())
	invoiceV1Group.PUT("/:invoice_id/payment_method", invoiceHandler.UpdatePaymentMethod(), params.Middleware.CustomerMustAuthorized())
	invoiceV1Group.GET("/reports", invoiceHandler.GetReport(), params.Middleware.CustomerMustAuthorized())
	invoiceV1Group.GET("/statuses", invoiceHandler.GetPaymentStatusList())
	invoiceV1Group.GET("/download/:invoice_id", invoiceHandler.DownloadInvoice(), params.Middleware.CustomerMustAuthorized())
}
