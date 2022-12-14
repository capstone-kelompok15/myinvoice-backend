package handler

import (
	"net/http"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/websocketutils"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type websocketHandler struct {
	log           *logrus.Entry
	validator     *validatorutils.Validator
	websocketPool *websocketutils.Pool
	upgrade       websocket.Upgrader
}

type WebsocketHandlerParams struct {
	Log           *logrus.Entry
	Validator     *validatorutils.Validator
	WebSocketPool *websocketutils.Pool
}

func NewWebsocketHandler(params *WebsocketHandlerParams) *websocketHandler {
	return &websocketHandler{
		log:           params.Log,
		validator:     params.Validator,
		websocketPool: params.WebSocketPool,
		upgrade: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}
}
