package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig.SMTPHost = viper.GetString("SMTP_HOST")
	AppConfig.SMTPPort = viper.GetString("SMTP_PORT")
	AppConfig.SMTPUser = viper.GetString("SMTP_USER")
	AppConfig.SMTPPassword = viper.GetString("SMTP_PASSWORD")
}
