package config

import (
	"api-gmr/env"
	"log"

	"github.com/joho/godotenv"
)

//App to hold common application configuration
type App struct {
	//AppName application name
	AppName string
	//AppPort port will be used for api sevice
	AppPort string
	//DbHost database host value
	DbHost string
	//DbPort database port value
	DbPort string
	//DbUser database user value
	DbUser string
	//DbPassword database password value
	DbPassword string
	//DbName database name value
	DbName string
	//DbDriver database driver. available drivers: mysql
	DbDriver string
	//JwtSecret secret key for signing jwt token
	JwtSecret string
	//Timezone timezone app
	TimeZone string
}

var app App

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Can't load .env file")
	}

	app = App{
		AppName:    env.GetValue("APP_NAME", "APP_GMR"),
		AppPort:    env.GetValue("APP_PORT", "8080"),
		DbHost:     env.GetValue("DB_HOST", "localhost"),
		DbPort:     env.GetValue("DB_PORT", "3306"),
		DbUser:     env.GetValue("DB_USER", "user"),
		DbPassword: env.GetValue("DB_PASSWORD", "password"),
		DbName:     env.GetValue("DB_NAME", "dbname"),
		DbDriver:   env.GetValue("DB_DRIVER", "mysql"),
		JwtSecret:  env.GetValue("JWT_SECRET", "jwt-secret-007"),
		TimeZone:   env.GetValue("TIMEZONE", "Asia/Jakarta"),
	}
}

//GetApp returns a new App instance
func GetApp() App {
	return app
}
