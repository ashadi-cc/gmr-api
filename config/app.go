package config

import (
	"api-gmr/env"
	"log"

	"github.com/joho/godotenv"
)

type App struct {
	AppName    string
	AppPort    string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	JwtSecret  string
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
		JwtSecret:  env.GetValue("JWT_SECRET", "jwt-secret-007"),
	}
}

func GetApp() App {
	return app
}
