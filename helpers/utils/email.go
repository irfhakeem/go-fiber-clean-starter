package utils

import (
	"log"

	"github.com/irfhakeem/go-fiber-clean-starter/config"
	"gopkg.in/gomail.v2"
)

func SendEmail(receiver string, body string, subject string) error {
	emailConf := config.EmailConf()

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", emailConf.AUTH_EMAIL)
	mailer.SetHeader("To", receiver)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		emailConf.HOST,
		emailConf.PORT,
		emailConf.AUTH_EMAIL,
		emailConf.AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	log.Println("Email sent!")
	return nil
}
