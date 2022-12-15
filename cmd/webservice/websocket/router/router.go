package router

import (
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/apiversioning"
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/websocket/handler"
	custommiddleware "github.com/capstone-kelompok15/myinvoice-backend/pkg/middleware/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/websocketutils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type WebSocketRouterParams struct {
	E             *echo.Echo
	Log           *logrus.Entry
	Validator     *validatorutils.Validator
	Middleware    custommiddleware.Middleware
	WebSocketPool *websocketutils.Pool
}

func InitWebSocketRouter(params *WebSocketRouterParams) {
	websocketHandler := handler.NewWebsocketHandler(&handler.WebsocketHandlerParams{
		Log:           params.Log,
		Validator:     params.Validator,
		WebSocketPool: params.WebSocketPool,
	})

	websocketV1Group := params.E.Group(apiversioning.APIVersionOne + "/ws")
	websocketV1Group.GET("/customers", websocketHandler.WebSocketClientCustomers(), params.Middleware.CustomerMustAuthorized())
	websocketV1Group.GET("/merchants", websocketHandler.WebSocketClientCustomers(), params.Middleware.AdminMustAuthorized())
}
