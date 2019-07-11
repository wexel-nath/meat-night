package config

import (
	"github.com/spf13/viper"
	"time"
)

func Configure() {

	// Service
	viper.Set("COMPANY_NAME", "Mateo Corporation")
	viper.Set("COMPANY_EMAIL", "mateocorp@getwexel.com")
	viper.Set("BASE_URL", "https://mateo-meat-night.herokuapp.com")
	viper.Set("DINNER_DAY", time.Wednesday)

	// Heroku Port
	viper.BindEnv("PORT")

	// Postgres URL
	viper.BindEnv("DATABASE_URL")

	// Mailgun
	viper.BindEnv("WEXEL_MAILGUN_DOMAIN")
	viper.BindEnv("WEXEL_MAILGUN_API_KEY")
	viper.BindEnv("WEXEL_MAILGUN_PUBLIC_KEY")
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
	return time.Weekday(viper.GetInt("DINNER_DAY"))
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
