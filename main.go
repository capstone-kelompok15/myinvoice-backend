package main

import (
	"log"

	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice"
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/logrusutils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := config.GetConfig("./.env")
	if err != nil {
		log.Fatal(err)
	}

	log := logrusutils.New()

	webservice.InitWebService(&webservice.WebServiceParams{
		Config: config,
		Log:    log,
	})
}
