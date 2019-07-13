package config

import (
	"time"

	"github.com/spf13/viper"
)

func Configure() {

	// Service
	viper.Set("COMPANY_NAME", "Mateo Corporation")
	viper.Set("COMPANY_EMAIL", "mateocorp@getwexel.com")
	viper.Set("DINNER_DAY", time.Wednesday)

	// Base url
	viper.SetDefault("BASE_URL", "http://localhost:4000")
	viper.BindEnv("BASE_URL")

	// Heroku Port
	viper.SetDefault("PORT", "4000")
	viper.BindEnv("PORT")

	// Postgres URL
	viper.SetDefault("DATABASE_URL", "host='localhost' port='5432' user='nathanw' password='bonnie' database='meat_night'")
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

func GetBaseURL() string {
	return viper.GetString("BASE_URL")
}

func GetDinnerDay() time.Weekday {
	return viper.Get("DINNER_DAY").(time.Weekday)
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
