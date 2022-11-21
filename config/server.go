package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
)

type Server struct {
	E    *echo.Echo
	Port int
}

func StartServer(param Server) error {
	errChan := make(chan error, 1)
	defer param.E.Shutdown(context.Background())

	go func() {
		if err := param.E.Start(fmt.Sprintf(":%d", param.Port)); err != nil {
			errChan <- err
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Println("[ERROR] Error while starting server: ", err.Error())
		log.Println("[INFO] Server closed gracefully")
		return err
	case <-signalChan:
		log.Println("[INFO] Server closed gracefully")
		return nil
	}
}
