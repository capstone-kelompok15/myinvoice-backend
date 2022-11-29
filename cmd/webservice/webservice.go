package webservice

import (
	"log"

	authrouter "github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/auth/router"
	pingrouter "github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/ping/router"
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	authrepository "github.com/capstone-kelompok15/myinvoice-backend/internal/auth/repository/impl"
	authservice "github.com/capstone-kelompok15/myinvoice-backend/internal/auth/service/impl"
	pingservice "github.com/capstone-kelompok15/myinvoice-backend/internal/ping/service/impl"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type WebServiceParams struct {
	Config *config.Config
	Log    *logrus.Logger
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

	validator, err := validatorutils.New()
	if err != nil {
		params.Log.Warningln("[ERROR] while creating the validator")
		return err
	}

	redis, err := config.InitRedis(&params.Config.RedisConfig)
	if err != nil {
		params.Log.Warningln("[ERROR] while creating the redis client")
		return err
	}

	mailgunClient := config.InitMailgun(&params.Config.Mailgun)

	pingService := pingservice.NewPingService(pingservice.Service{})
	pingrouter.InitPingRouter(pingrouter.RouterParams{
		E:           e,
		PingService: pingService,
	})

	customerRepository := authrepository.NewCustomerRepository(&authrepository.CustomerRepositoryParams{
		DB: db,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "customer",
			"layer":  "repository",
		}),
	})

	customerService := authservice.NewCustomerService(&authservice.CustomerServiceParams{
		Repo:    customerRepository,
		Redis:   redis,
		Mailgun: mailgunClient,
		Config:  params.Config,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "customer",
			"layer":  "service",
		}),
	})

	authrouter.InitCustomerRouter(&authrouter.RouterParams{
		E:         e,
		Service:   customerService,
		Validator: validator,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "customer",
			"layer":  "handler",
		}),
	})

	err = config.StartServer(config.Server{
		E:    e,
		Port: params.Config.Server.Port,
	})

	if err != nil {
		params.Log.Warningln("[ERROR] while starting the server")
		return err
	}

	return nil
}
