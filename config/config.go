package config

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

type Config struct {
	AppName     string
	AppPort     int
	LogLevel    string
	Environment string
	JWTSecret   string
	MYSQL_USER  string
	MYSQL_PORT  string
	MYSQL_PASS  string
	MYSQL_ADDR  string
}

func Init() *Config {
	defaultEnv := ".env"

	if err := gotenv.Load(defaultEnv); err != nil {
		log.Fatal("failed load .env")
	}

	log.SetOutput(os.Stdout)

	appConfig := &Config{
		AppName:     GetString("APP_NAME"),
		AppPort:     GetInt("APP_PORT"),
		LogLevel:    GetString("LOG_LEVEL"),
		Environment: GetString("ENVIRONMENT"),
		JWTSecret:   GetString("JWT_SECRET"),
		MYSQL_USER:  GetString("mysql_user"),
		MYSQL_PORT:  GetString("mysql_port"),
		MYSQL_PASS:  GetString("mysql_pass"),
		MYSQL_ADDR:  GetString("mysql_addr"),
	}

	return appConfig
}
