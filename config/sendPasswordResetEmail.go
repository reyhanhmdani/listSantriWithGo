package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"os"
)

func SendPasswordResetEmail(email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "0yt0premium0@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset")
	m.SetBody("text/html", "To reset your password, use this token: "+token)

	SmtpUser := os.Getenv("SMTP_USER")
	SmtpPassword := os.Getenv("SMTP_PASS")

	//log.Printf("smtp Password: %s", SmtpPassword)
	////log.Printf("DB Database: %s", dbDatabase)

	d := gomail.NewDialer("smtp.mailtrap.io", 2525, SmtpUser, SmtpPassword)

	// Kirim email
	if err := d.DialAndSend(m); err != nil {
		logrus.Error("error di Dial And Send ", err)
		return err
	}

	return nil
}
