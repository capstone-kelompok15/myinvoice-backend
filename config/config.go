package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database
	JWTConfig
	Server
	Cloudinary
}

var config *Config

func initConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("[INFO] The .env file doesn't exist")
		log.Println("[INFO] Program will load environment variable value")
	}

	portStr := os.Getenv("APP_PORT")
	if portStr == "" {
		log.Fatal("[ERROR] Error while init web server, app port cant be empty")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("[ERROR] Error while init web server, app port must be a number")
	}

	dbPortStr := os.Getenv("DB_PORT")
	if portStr == "" {
		log.Fatal("[ERROR] Error while init database, database port cant be empty")
	}

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Fatal("[ERROR] Error while init database, database port must be a number")
	}

	config = &Config{
		Database: Database{
			Username:                     os.Getenv("DB_USERNAME"),
			Password:                     os.Getenv("DB_PASSWORD"),
			Hostname:                     os.Getenv("DB_HOSTNAME"),
			Port:                         dbPort,
			DatabaseName:                 os.Getenv("DB_NAME"),
			RelationalDatabaseDriverName: os.Getenv("DB_DRIVER_NAME"),
		},
		JWTConfig: JWTConfig{
			JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
		},
		Server: Server{
			Port: port,
		},
		Cloudinary: Cloudinary{
			APIKey:    os.Getenv("API_KEY"),
			APISecret: os.Getenv("API_SECRET"),
			CloudName: os.Getenv("CLOUD_NAME"),
		},
	}

	if config.Database.Username == "" {
		log.Fatal("[ERROR] Error while init database, database name cant be empty")
	}

	if config.Database.Password == "" {
		log.Fatal("[ERROR] Error while init database, database password cant be empty")
	}

	if config.Database.Hostname == "" {
		log.Fatal("[ERROR] Error while init database, database hostname cant be empty")
	}

	if config.Database.DatabaseName == "" {
		log.Fatal("[ERROR] Error while init database, database name cant be empty")
	}

	if config.Cloudinary.APISecret == "" {
		log.Fatal("[ERROR] Error while init cloudinary, api secret cant be empty")
	}

	if config.Cloudinary.APIKey == "" {
		log.Fatal("[ERROR] Error while init cloudinary, api key cant be empty")
	}

	if config.Cloudinary.CloudName == "" {
		log.Fatal("[ERROR] Error while init cloudinary, cloud name cant be empty")
	}

	if config.JWTConfig.JWTSecretKey == "" {
		log.Fatal("[ERROR] Error while init jwt config, jwt secret key cant be empty")
	}

	if config.JWTConfig.JWTSecretKey == "" {
		log.Fatal("[ERROR] Error while init jwt config, jwt secret key cant be empty")
	}
}

func GetConfig() *Config {
	if config == nil {
		initConfig()
	}
	return config
}
