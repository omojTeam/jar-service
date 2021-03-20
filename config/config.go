package config

import (
	"fmt"

	v "github.com/spf13/viper"
)

type Config struct {
	DbAccess string
	Port     string
}

var Cfg Config

func LoadConfig() error {
	v.AutomaticEnv()
	Cfg.Port = v.GetString("PORT")

	if v.GetString("ENV") == "dev" {
		Cfg.DbAccess = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			v.GetString("HOST"), v.GetString("DBPORT"), v.GetString("USER"), v.GetString("PASSWORD"), v.GetString("DBNAME"))
	} else {
		Cfg.DbAccess = v.GetString("DATABASE_URL")
	}

	return nil
}
