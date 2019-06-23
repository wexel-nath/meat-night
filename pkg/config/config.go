package config

import (
	"github.com/spf13/viper"
)

func Configure() {

	// Service
	viper.Set("COMPANY_NAME", "Mateo Corporation")
	viper.Set("COMPANY_EMAIL", "nathanwelch_@hotmail.com")

	// Heroku Port
	viper.BindEnv("PORT")

	// Postgres URL
	viper.BindEnv("DATABASE_URL")

	// Mailgun
	viper.BindEnv("MAILGUN_DOMAIN")
	viper.BindEnv("MAILGUN_API_KEY")
	viper.BindEnv("MAILGUN_PUBLIC_KEY")
}

func GetPort() string {
	return viper.GetString("PORT")
}

func GetDatabaseURL() string {
	return viper.GetString("DATABASE_URL")
}

func GetMailgunDomain() string {
	return viper.GetString("MAILGUN_DOMAIN")
}

func GetMailgunApiKey() string {
	return viper.GetString("MAILGUN_API_KEY")
}

func GetMailgunPublicKey() string {
	return viper.GetString("MAILGUN_PUBLIC_KEY")
}

func GetCompanyName() string {
	return viper.GetString("COMPANY_NAME")
}

func GetCompanyEmail() string {
	return viper.GetString("COMPANY_EMAIL")
}
