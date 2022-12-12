package webservice

import (
	"strings"

	authrouter "github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/auth/router"
	bankrouter "github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/bank/router"
	customerrouter "github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/customer/router"
	invoicerouter "github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/invoice/router"
	merchantrouter "github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/merchant/router"
	pingrouter "github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/ping/router"
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	authrepository "github.com/capstone-kelompok15/myinvoice-backend/internal/auth/repository/impl"
	authservice "github.com/capstone-kelompok15/myinvoice-backend/internal/auth/service/impl"
	bankrepository "github.com/capstone-kelompok15/myinvoice-backend/internal/bank/repository/impl"
	bankservice "github.com/capstone-kelompok15/myinvoice-backend/internal/bank/service/impl"
	customerrepository "github.com/capstone-kelompok15/myinvoice-backend/internal/customer/repository/impl"
	customerservice "github.com/capstone-kelompok15/myinvoice-backend/internal/customer/service/impl"
	invoicerepository "github.com/capstone-kelompok15/myinvoice-backend/internal/invoice/repository/impl"
	invoiceservice "github.com/capstone-kelompok15/myinvoice-backend/internal/invoice/service/impl"
	merchantrepository "github.com/capstone-kelompok15/myinvoice-backend/internal/merchant/repository/impl"
	merchantservice "github.com/capstone-kelompok15/myinvoice-backend/internal/merchant/service/impl"
	notificationrepository "github.com/capstone-kelompok15/myinvoice-backend/internal/notification/repository/impl"
	pingservice "github.com/capstone-kelompok15/myinvoice-backend/internal/ping/service/impl"
	customrepositorymiddleware "github.com/capstone-kelompok15/myinvoice-backend/pkg/middleware/repository/impl"
	customservicemiddleware "github.com/capstone-kelompok15/myinvoice-backend/pkg/middleware/service/impl"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type WebServiceParams struct {
	Config *config.Config
	Log    *logrus.Logger
}

func InitWebService(params *WebServiceParams) error {
	db, err := config.GetDatabaseConn(&params.Config.Database)
	if err != nil {
		params.Log.Warningln("[ERROR] while get database connection:", err.Error())
		return err
	}

	defer func() error {
		err := config.CloseDatabaseConnection(db)
		if err != nil {
			params.Log.Warningln("[ERROR] while close database connection:", err.Error())
			return err
		}
		params.Log.Warningln("[INFO] Database connection closed gracefully:", err.Error())
		return nil
	}()

	whiteListAllowOrigin := strings.Split(params.Config.Server.WhiteListAllowOrigin, ",")

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     whiteListAllowOrigin,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	validator, err := validatorutils.New()
	if err != nil {
		params.Log.Warningln("[ERROR] while creating the validator:", err.Error())
		return err
	}

	redis, err := config.InitRedis(&params.Config.RedisConfig)
	if err != nil {
		params.Log.Warningln("[ERROR] while creating the redis client:", err.Error())
		return err
	}

	mailgunClient := config.InitMailgun(&params.Config.Mailgun)

	cloudinary, err := config.GetCloudinaryConn(&params.Config.Cloudinary)
	if err != nil {
		params.Log.Warningln("[ERROR] while creating the cloudinary client:", err.Error())
		return err
	}

	// Middleware
	repositoryMiddleware := customrepositorymiddleware.NewRepositoryMiddleware(&customrepositorymiddleware.MiddlewareRepositoryParams{
		DB: db,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "middleware",
			"layer":  "repository",
		}),
	})

	middleware := customservicemiddleware.NewServiceMiddleware(&customservicemiddleware.MiddlewareParams{
		Config:         params.Config,
		Redis:          redis,
		MiddlewareRepo: repositoryMiddleware,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "middleware",
			"layer":  "service",
		}),
	})

	// Ping
	pingService := pingservice.NewPingService(pingservice.Service{})
	pingrouter.InitPingRouter(pingrouter.RouterParams{
		E:           e,
		PingService: pingService,
	})

	// Notification
	notificationRepository := notificationrepository.NewNotificationRepository(&notificationrepository.NotificationRepositoryParams{
		DB: db,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "notification",
			"layer":  "repository",
		}),
	})

	// Auth
	authRepository := authrepository.NewAuthRepository(&authrepository.AuthRepositoryParams{
		DB: db,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "auth",
			"layer":  "repository",
		}),
	})

	authService := authservice.NewAuthService(&authservice.AuthServiceParams{
		Repo:    authRepository,
		Redis:   redis,
		Mailgun: mailgunClient,
		Config:  params.Config,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "auth",
			"layer":  "service",
		}),
	})

	authrouter.InitAuthRouter(&authrouter.RouterParams{
		E:         e,
		Service:   authService,
		Validator: validator,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "auth",
			"layer":  "handler",
		}),
	})

	// Customer
	customerRepository := customerrepository.NewCustomerRepository(&customerrepository.CustomerRepositoryParams{
		DB: db,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "customer",
			"layer":  "repository",
		}),
	})

	customerService := customerservice.NewCustomerService(&customerservice.CustomerServiceParams{
		RepoNotif:  notificationRepository,
		Repo:       customerRepository,
		Redis:      redis,
		Mailgun:    mailgunClient,
		Config:     params.Config,
		Cloudinary: cloudinary,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "customer",
			"layer":  "service",
		}),
	})

	customerrouter.InitCustomerRouter(&customerrouter.RouterParams{
		E:          e,
		Validator:  validator,
		Service:    customerService,
		Middleware: middleware,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "auth",
			"layer":  "handler",
		}),
	})

	// Bank
	bankRepository := bankrepository.NewBankRepository(&bankrepository.BankRepositoryParams{
		DB: db,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "bank",
			"layer":  "repository",
		}),
	})

	bankService := bankservice.NewBankService(&bankservice.BankServiceParams{
		Repo:   bankRepository,
		Config: params.Config,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "bank",
			"layer":  "service",
		}),
	})

	bankrouter.InitBankRouter(&bankrouter.BankRouterParams{
		E:         e,
		Service:   bankService,
		Validator: validator,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "bank",
			"layer":  "handler",
		}),
	})

	// Merchant
	merchantRepository := merchantrepository.NewBankRepository(&merchantrepository.MerchantRepositoryParams{
		DB: db,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "merchant",
			"layer":  "repository",
		}),
	})

	merchantService := merchantservice.NewMerchantService(&merchantservice.MerchantServiceParams{
		RepoNotif:  notificationRepository,
		Repo:       merchantRepository,
		Config:     params.Config,
		Cloudinary: cloudinary,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "bank",
			"layer":  "service",
		}),
	})

	merchantrouter.InitMerchantRouter(&merchantrouter.MerchantRouterParams{
		E:          e,
		Validator:  validator,
		Service:    merchantService,
		Middleware: middleware,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "merchant",
			"layer":  "handler",
		}),
	})

	// Invoices
	invoiceRepository := invoicerepository.NewInvoiceRepository(&invoicerepository.InvoiceRepositoryParams{
		DB: db,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "invoice",
			"layer":  "repository",
		}),
	})

	invoiceService := invoiceservice.NewInvoiceService(&invoiceservice.InvoiceService{
		Repo:    invoiceRepository,
		Config:  params.Config,
		Mailgun: mailgunClient,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "invoice",
			"layer":  "service",
		}),
	})

	invoicerouter.InitInvoiceRouter(&invoicerouter.InvoiceRouterParams{
		E:          e,
		Validator:  validator,
		Service:    invoiceService,
		Middleware: middleware,
		Log: params.Log.WithFields(logrus.Fields{
			"domain": "invoice",
			"layer":  "handler",
		}),
	})

	err = config.StartServer(config.Server{
		E:    e,
		Port: params.Config.Server.Port,
	})

	if err != nil {
		params.Log.Warningln("[ERROR] while starting the server:", err.Error())
		return err
	}

	return nil
}
