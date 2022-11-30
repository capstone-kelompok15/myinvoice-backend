package config

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	Username                     string `validate:"required"`
	Password                     string `validate:"required"`
	Hostname                     string `validate:"required"`
	Port                         string `validate:"required"`
	DatabaseName                 string `validate:"required"`
	RelationalDatabaseDriverName string `validate:"required"`
}

var db *sqlx.DB

func initDatabase(params *Database) error {
	port, err := strconv.Atoi(params.Port)
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		params.Username, params.Password,
		params.Hostname, port,
		params.DatabaseName,
	)

	var secondAttempt int64 = 1

	for {
		if db, err = sqlx.Open(params.RelationalDatabaseDriverName, dsn); err == nil {
			break
		}

		if err = db.Ping(); err == nil {
			break
		}

		log.Println("[HANDLER ERROR] Cant establish the database connection, trying in 1 second")
		time.Sleep(time.Duration(secondAttempt) * time.Second)
		secondAttempt++
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	log.Println("[INFO] Successfully establishing database connection")
	return nil
}

func GetDatabaseConn(params *Database) (*sqlx.DB, error) {
	var err error

	if db == nil {
		err = initDatabase(params)
		if err != nil {
			log.Println("[ERROR] While get database connection, init database:", err.Error())
			return nil, err
		}
	}

	return db, nil
}

func CloseDatabaseConnection(db *sqlx.DB) error {
	err := db.Close()
	return err
}
