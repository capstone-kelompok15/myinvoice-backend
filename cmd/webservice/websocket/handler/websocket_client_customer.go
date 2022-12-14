package handler

import (
	"log"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/authutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/websocketutils"
	"github.com/labstack/echo/v4"
)

func (h *websocketHandler) WebSocketClientCustomers() echo.HandlerFunc {
	return func(c echo.Context) error {
		userCtx := authutils.CustomerFromRequestContext(c)
		if userCtx == nil {
			log.Println("[WebSocketClient] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: customerrors.ErrInternalServer,
			})
		}

		wsConn, err := h.upgrade.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			log.Println(err.Error())
			return nil
		}

		client := &websocketutils.Client{
			CustomerID: userCtx.ID,
			Conn:       wsConn,
			Pool:       h.websocketPool,
		}
		defer client.Close()

		client.Pool.Register <- client
		client.Conn.WriteMessage(1, []byte("Connected to websocket"))
		client.Read()
		return nil
	}
}
