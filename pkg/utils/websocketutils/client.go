package websocketutils

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	CustomerID int
	MerchantID int
	Conn       *websocket.Conn
	Pool       *Pool
}

func (c *Client) Read() {
	for {
		var pJson Message

		err := c.Conn.ReadJSON(&pJson)
		if err != nil {
			log.Println(err)
			return
		}

		if pJson.Disconnect {
			c.Close()
		}
	}
}

func (c *Client) Close() {
	c.Conn.WriteMessage(1, []byte("Disconnected from websocket"))
	c.Pool.Unregister <- c
	c.Conn.Close()
}
