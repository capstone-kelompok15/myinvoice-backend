package webservice

import (
	"log"

	pingrouter "github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/ping/router"
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	pingservice "github.com/capstone-kelompok15/myinvoice-backend/internal/ping/service/impl"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type WebServiceParams struct {
	Config *config.Config
}

func InitWebService(params *WebServiceParams) error {
	db, err := config.GetDatabaseConn(&params.Config.Database)
	if err != nil {
		log.Println("[ERROR] while get database connection")
		return err
	}

	defer func() error {
		err := config.CloseDatabaseConnection(db)
		if err != nil {
			log.Println("[ERROR] while close database connection")
			return err
		}
		log.Println("[INFO] Database connection closed gracefully")
		return nil
	}()

	// cloudinary, err := config.GetCloudinaryConn(&params.Config.Cloudinary)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	e := echo.New()

	pingService := pingservice.NewPingService(pingservice.Service{})
	pingrouter.InitPingRouter(pingrouter.RouterParams{
		E:           e,
		PingService: pingService,
	})

	err = config.StartServer(config.Server{
		E:    e,
		Port: params.Config.Server.Port,
	})

	if err != nil {
		log.Println("[ERROR] while starting the server")
		return err
	}

	return nil
}
