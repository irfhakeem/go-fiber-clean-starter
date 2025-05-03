package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EmailConfig struct {
	HOST          string
	PORT          int
	SENDER_NAME   string
	AUTH_EMAIL    string
	AUTH_PASSWORD string
}

func EmailConf() EmailConfig {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file: " + err.Error())
		}
	}

	port := func() int {
		port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
		if err != nil {
			log.Fatalf("Invalid SMTP_PORT: %v", err)
		}
		return port
	}()

	email := EmailConfig{
		HOST:          os.Getenv("SMTP_HOST"),
		PORT:          port,
		SENDER_NAME:   os.Getenv("SMTP_SENDER_NAME"),
		AUTH_EMAIL:    os.Getenv("SMTP_AUTH_EMAIL"),
		AUTH_PASSWORD: os.Getenv("SMTP_AUTH_PASSWORD"),
	}

	return email
}
