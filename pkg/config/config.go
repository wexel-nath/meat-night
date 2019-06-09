package config

import (
	"github.com/spf13/viper"
)

func Configure() {

	// Heroku
	viper.BindEnv("DATABASE_URL")
	viper.BindEnv("PORT")
}

func GetDatabaseURL() string {
	return viper.GetString("DATABASE_URL")
}

func GetPort() string {
	return viper.GetString("PORT")
}
