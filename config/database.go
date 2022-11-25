package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Database struct {
	Username                     string
	Password                     string
	Hostname                     string
	Port                         int
	DatabaseName                 string
	RelationalDatabaseDriverName string
}

var db *sql.DB

func initDatabase(params *Database) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		params.Username, params.Password,
		params.Hostname, params.Port,
		params.DatabaseName,
	)

	var err error
	var secondAttempt int64 = 1

	for {
		if db, err = sql.Open(params.RelationalDatabaseDriverName, dsn); err == nil {
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

func GetDatabaseConn(params *Database) (*sql.DB, error) {
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

func CloseDatabaseConnection(db *sql.DB) error {
	err := db.Close()
	return err
}
