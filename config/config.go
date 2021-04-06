package config

import (
	"fmt"

	v "github.com/spf13/viper"
)

type Config struct {
	DbAccess      string
	Port          string
	EmailHost     string
	EmailPort     int
	EmailUsername string
	EmailPassword string
}

var Cfg Config

func LoadConfig() error {
	v.AutomaticEnv()
	Cfg.Port = v.GetString("PORT")
	Cfg.EmailHost = v.GetString("EMAIL_HOST")
	Cfg.EmailPort = v.GetInt("EMAIL_PORT")
	Cfg.EmailUsername = v.GetString("EMAIL_LOGIN")
	Cfg.EmailPassword = v.GetString("EMAIL_PASSWORD")

	if v.GetString("ENV") == "dev" {
		Cfg.DbAccess = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			v.GetString("HOST"), v.GetString("DBPORT"), v.GetString("USER"), v.GetString("PASSWORD"), v.GetString("DBNAME"))
	} else {
		Cfg.DbAccess = v.GetString("DATABASE_URL")
	}

	return nil
}
