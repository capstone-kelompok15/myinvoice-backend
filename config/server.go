package config

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/labstack/echo/v4"
)

type Server struct {
	E                    *echo.Echo
	Port                 string `validate:"required"`
	Environment          string `validate:"oneof='dev' 'prod'"`
	WhiteListAllowOrigin string
}

func StartServer(param Server) error {
	port, err := strconv.Atoi(param.Port)
	if err != nil {
		log.Println("[ERROR] Error while convert port to number:", err.Error())
		return err
	}

	errChan := make(chan error, 1)

	go func() {
		if err := param.E.Start(fmt.Sprintf(":%d", port)); err != nil {
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
		return nil
	}
}
