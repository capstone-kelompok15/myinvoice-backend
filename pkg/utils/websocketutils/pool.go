package websocketutils

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
)

type Pool struct {
	Log        *logrus.Entry
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Message    chan *Message
}

type PoolParams struct {
	Log *logrus.Entry
}

func NewPool(params *PoolParams) *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Message:    make(chan *Message),
		Clients:    make(map[*Client]bool),
		Log:        params.Log,
	}
}

func (pool *Pool) Start() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-signalChan:
			pool.Close()
			return
		case client := <-pool.Register:
			pool.Clients[client] = true
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
		case message := <-pool.Message:
			groupSync := sync.WaitGroup{}
			// Customer
			groupSync.Add(2)

			go func() {
				defer groupSync.Done()
				for client := range pool.Clients {
					if client.CustomerID == message.SendToCustomer {
						if err := client.Conn.WriteJSON(message); err != nil {
							pool.Log.Warningln("[WebSocket] Error on send to the customer:", err.Error())
							return
						}
					}
				}
			}()

			// Merchant
			go func() {
				defer groupSync.Done()
				for client := range pool.Clients {
					if client.MerchantID == message.SendToMerchant {
						if err := client.Conn.WriteJSON(message); err != nil {
							pool.Log.Warningln("[WebSocket] Error on send to the merchant:", err.Error())
							return
						}
					}
				}
			}()

			groupSync.Wait()
		}
	}
}

func (pool *Pool) Close() {
	for client := range pool.Clients {
		client.Close()
	}
	close(pool.Message)
	close(pool.Unregister)
	close(pool.Register)
	pool.Log.Infoln("[INFO] WebSocket connection closed gracefully")
}
