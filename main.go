package main

import (
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice"
	"github.com/capstone-kelompok15/myinvoice-backend/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := config.GetConfig()

	webservice.InitWebService(&webservice.WebServiceParams{
		Config: config,
	})
}
